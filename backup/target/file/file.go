package file

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/sikalabs/tergum/utils/file_utils"
)

type FileTarget struct {
	Dir    string `yaml:"Dir"`
	Prefix string `yaml:"Prefix"`
	Suffix string `yaml:"Suffix"`
}

func (t FileTarget) Validate() error {
	if t.Dir == "" {
		return errors.New("file requires dir")
	}
	if t.Prefix == "" {
		return errors.New("file requires prefix")
	}
	if t.Suffix == "" {
		return errors.New("file requires suffix")
	}
	return nil
}

func (t FileTarget) Save(data io.ReadSeeker) error {
	err := os.MkdirAll(t.Dir, 0755)
	if err != nil {
		return err
	}

	name := file_utils.GetFileName(filepath.Join(t.Dir, t.Prefix), t.Suffix)

	if err != nil {
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, data)
	return err

}
