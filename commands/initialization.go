package commands

import (
	"strings"

	"github.com/chainpusher/chainpusher/config"
	"github.com/sirupsen/logrus"
)

func SetupLogger(c *config.Config) {
	switch strings.ToUpper(c.Logger.Level) {
	case "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Info("Logger level set to ", logrus.GetLevel())

}
