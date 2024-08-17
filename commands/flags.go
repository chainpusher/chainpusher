package commands

import (
	"errors"
	"github.com/chainpusher/chainpusher/config"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func SetupConfigPathFlag(cmd *cobra.Command) {
	cmd.Flags().String("config", "p", "Path to the config file")
}

func GetConfig(cmd *cobra.Command) (*config.Config, error) {
	var p string
	var err error
	var wd string
	var home string
	var directories []string

	if p, err = cmd.Flags().GetString("config"); err != nil {
		p = "./config.yaml"
	}

	if wd, err = os.Getwd(); err == nil {
		directories = append(directories, path.Join(wd, "config.yaml"))
	}

	if home, err = os.UserHomeDir(); err != nil {
		directories = append(directories, path.Join(home, "chainpusher"))
	}

	if path.IsAbs(p) {
		return config.ParseConfigFromYaml(p)
	}

	for _, d := range directories {
		if _, err = os.Stat(d); err == nil {
			return config.ParseConfigFromYaml(d)
		}
	}

	return nil, errors.New("config file not found")
}
