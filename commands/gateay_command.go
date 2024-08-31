package commands

import (
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/module/datasource"
	"github.com/chainpusher/chainpusher/module/task"
	"github.com/chainpusher/chainpusher/trade/infrastructure/factory"
	cron2 "github.com/chainpusher/chainpusher/trade/interfaces/price/cron"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func GatewayCommand(cmd *cobra.Command, args []string) {
	var c *cron.Cron
	var cfg *config.Config
	var j cron.Job
	var db *gorm.DB
	var t task.Task
	var err error
	if cfg, err = GetConfig(cmd); err != nil {
		panic(err)
	}

	ds := cfg.Gateway.Datasource["default"]
	if db, err = datasource.GetGormDatabase(&ds); err != nil {
		panic(err)
	}

	j = cron2.NewTradePriceJob(factory.NewFactory(db))
	if c, err = cron2.NewCron(j); err != nil {
		panic(err)
	}

	t = cron2.NewTask(c)
	m := task.NewManager(t)
	m.Start()

	m.Wait()
}

func NewGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Starts the gateway",
		Run:   GatewayCommand,
	}

	SetupConfigPathFlag(cmd)

	return cmd
}
