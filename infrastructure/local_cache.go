package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"
)

var KEY_PREFIX = "chainpusher"

type Cache []byte

func GetKey(key string) (Cache, error) {
	filename := fmt.Sprintf("%s_%s.cached", KEY_PREFIX, key)
	path := filepath.Join(os.TempDir(), filename)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func SetKey(key string, value []byte) error {
	filename := fmt.Sprintf("%s_%s.cached", KEY_PREFIX, key)
	path := filepath.Join(os.TempDir(), filename)
	err := os.WriteFile(path, value, 0644)

	if err != nil {
		return err
	}
	return nil
}

func (c Cache) String() string {
	return string(c)
}
