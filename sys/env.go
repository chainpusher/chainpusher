package sys

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var DotEnvFileLoaded bool = false

func LoadDotEnv() error {
	paths := []string{"./", ".."}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	loadPaths := make([]string, 0)

	for _, path := range paths {
		path = filepath.Join(wd, path, ".env")
		_, err := os.Stat(path)

		logrus.Info("Loading .env file from: ", path)

		if os.IsNotExist(err) {
			continue
		}

		loadPaths = append(loadPaths, path)
	}

	if len(loadPaths) == 0 {
		return errors.New("no .env file found")
	}

	err = godotenv.Load(loadPaths...)

	return err
}

func GetEnv(key string) (string, error) {
	if !DotEnvFileLoaded {
		if err := LoadDotEnv(); err != nil {
			return "", err
		}
	}
	return os.Getenv(key), nil
}
