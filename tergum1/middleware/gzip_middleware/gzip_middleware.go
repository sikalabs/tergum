package gzip_middleware

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/sikalabs/tergum/tergum1/utils/file_utils"
)

type File struct {
	Dir    string
	Prefix string
	Suffix string
}

func ValidateFile(config File) error {
	if config.Dir == "" {
		return errors.New("file requires dir")
	}
	if config.Prefix == "" {
		return errors.New("file requires prefix")
	}
	if config.Suffix == "" {
		return errors.New("file requires suffix")
	}
	return nil
}

func SaveFile(config File, data []byte) error {
	err := os.MkdirAll(config.Dir, 0755)
	if err != nil {
		return err
	}
	name := file_utils.GetFileName(config.Prefix, config.Suffix)
	err = ioutil.WriteFile(path.Join(config.Dir, name), data, 0644)
	return err
}
