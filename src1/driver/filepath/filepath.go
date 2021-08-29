package filepath

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilePath struct {
	Path string
}

func ValidateFilePath(config FilePath) error {
	if config.Path == "" {
		return errors.New("filepath requires path")
	}
	return nil
}

func SaveFilePath(config FilePath, data []byte) error {
	dir := filepath.Dir(config.Path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(config.Path, data, 0644)
	return err
}
