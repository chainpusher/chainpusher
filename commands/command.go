package commands

import (
	"fmt"
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/interfaces/facade/impl"
	"github.com/chainpusher/chainpusher/interfaces/web"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	monitor2 "github.com/chainpusher/chainpusher/monitor"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/chainpusher/chainpusher/config"
	"github.com/spf13/cobra"
)

type MonitorCommandOptions struct {

	// These
	Listeners []service.BlockListener

	Movement monitor2.Movement

	Listener Listener
}

type CommandRunner struct {
	ctx *monitor2.Ctx

	clients *socket.Clients

	tasks *sys.TaskManager
}

func (r *CommandRunner) Run() {
	defer r.recover()

	r.tasks.Start()

	done := make(chan bool, 1)
	go r.signal(done)
	<-done
}

func (r *CommandRunner) signal(done chan bool) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	aSignal := <-signals
	logrus.Debugf("Signal received: %s", aSignal)
	done <- true
}

func (r *CommandRunner) recover() {
	if r := recover(); r == nil {
		return
	}

	logrus.Warnf("Recovered from a panic: %v", r)
	r.printStackTrace()
}

func (r *CommandRunner) printStackTrace() {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, 2*len(buf))
	}
	fmt.Printf("Stack trace: %s\n", buf)
}

func NewCommandRunner(cmd *cobra.Command, options MonitorCommandOptions) *CommandRunner {
	ctx := createCtx(cmd, options)
	monitor := NewMonitorCommand(ctx)
	clients := socket.NewClients()
	svc := application.NewTinyBlockService(clients)
	facade := impl.NewTinyBlockServiceFacade(svc)
	processor := web.NewJsonRpcMessageProcess(facade)
	server := web.NewServerTask("localhost", 8080, processor, clients)

	tm := sys.NewTaskManager()
	tm.Add(server)
	tm.Add(monitor)

	return &CommandRunner{
		ctx:     ctx,
		clients: clients,
		tasks:   tm,
	}
}

func createCtx(cmd *cobra.Command, options MonitorCommandOptions) *monitor2.Ctx {

	var cfg *config.Config
	p, err := cmd.Flags().GetString("config")

	if err != nil {
		cfg = &config.Config{}
	} else {
		cfg, err = config.ParseConfigFromYaml(p)
		if err != nil {
			cfg = &config.Config{}
			logrus.Errorf("failed to parse config: %v", err)
		}
	}

	isTesting, err := cmd.Flags().GetBool("test")
	if err == nil {
		cfg.IsTesting = isTesting
	}

	cfg.BlockLoggingFile, _ = cmd.Flags().GetString("block-file")

	SetupLogger(cfg)

	ctx := monitor2.Ctx{
		Config:    cfg,
		Listeners: options.Listeners,
	}

	ctx.Movement = options.Movement

	if options.Listener != nil {
		options.Listener.ConfigLoaded(&ctx)
	}

	return &ctx
}

func NewMonitorCobraCommand(options MonitorCommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor blockchain data",
		Run: func(cmd *cobra.Command, _ []string) {
			runner := NewCommandRunner(cmd, options)
			runner.Run()
		},
	}

	cmd.PersistentFlags().StringP("block-file", "b", "", "File to write raw blockchain data to")
	cmd.PersistentFlags().StringP("trx-file", "t", "", "File to write transactions to")
	cmd.PersistentFlags().BoolP("test", "x", false, "Test mode")

	return cmd
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chainpusher",
		Short: "A CLI tool for pushing blockchain data",
		Long: "Chainpusher is a CLI tool for pushing blockchain data to a remote server. " +
			"Chainpusher can also monitor blockchain data and push it to a remote server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().String(
		"config",
		"c",
		"config file (default is $HOME/.chainpusher.yaml)",
	)

	return cmd
}

func RunCommand() {

	RunCommandWithOptions(MonitorCommandOptions{
		Listeners: []service.BlockListener{},
	})
}

func RunCommandWithOptions(options MonitorCommandOptions) {

	rootCmd := NewRootCommand()
	monitorCmd := NewMonitorCobraCommand(options)

	rootCmd.AddCommand(monitorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
