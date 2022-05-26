package single_file

import (
	"fmt"
	"io"
	"os"
)

type SingleFileSource struct {
	Path string `yaml:"Path"`
}

func (s SingleFileSource) Validate() error {
	if s.Path == "" {
		return fmt.Errorf("SingleFileSource need to have a FilePath")
	}
	return nil
}

func (s SingleFileSource) Backup() (io.ReadSeeker, string, error) {
	out, err := os.Open(s.Path)
	return out, "", err
}
