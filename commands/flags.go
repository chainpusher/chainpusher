package commands

import (
	"errors"
	"os"
	"path"

	"github.com/chainpusher/chainpusher/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
		p = "./config.yml"
	}

	if wd, err = os.Getwd(); err == nil {
		directories = append(directories, wd)
	}

	if home, err = os.UserHomeDir(); err != nil {
		directories = append(directories, path.Join(home, "chainpusher"))
	}

	if path.IsAbs(p) {
		return config.ParseConfigFromYaml(p)
	}

	for _, d := range directories {
		f := path.Join(d, p)
		if _, err = os.Stat(f); err == nil {
			return config.ParseConfigFromYaml(f)
		} else {
			logrus.Infof("Load config file (%v) fail %v", f, err)
		}
	}

	return nil, errors.New("config file not found")
}
