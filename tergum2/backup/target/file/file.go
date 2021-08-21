package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/sikalabs/tergum/tergum1/utils/file_utils"
)

type FileTarget struct {
	Dir    string
	Prefix string
	Suffix string
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

func (t FileTarget) Save(data []byte) error {
	err := os.MkdirAll(t.Dir, 0755)
	if err != nil {
		return err
	}
	name := file_utils.GetFileName(t.Prefix, t.Suffix)
	err = ioutil.WriteFile(path.Join(t.Dir, name), data, 0644)
	return err
}
