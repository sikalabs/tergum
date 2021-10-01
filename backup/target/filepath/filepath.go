package filepath

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

type FilePathTarget struct {
	Path string `yaml:"Path"`
}

func (t FilePathTarget) Validate() error {
	if t.Path == "" {
		return errors.New("FilePathTarget requires Path")
	}
	return nil
}

func (t FilePathTarget) Save(data io.Reader) error {
	dir := filepath.Dir(t.Path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(t.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, data)
	return err
}
