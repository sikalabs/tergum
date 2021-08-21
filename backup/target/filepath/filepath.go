package filepath

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilePathTarget struct {
	Path string
}

func (t FilePathTarget) Validate() error {
	if t.Path == "" {
		return errors.New("FilePathTarget requires Path")
	}
	return nil
}

func (t FilePathTarget) Save(data []byte) error {
	dir := filepath.Dir(t.Path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(t.Path, data, 0644)
	return err
}
