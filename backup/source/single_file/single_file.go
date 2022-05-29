package single_file

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
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

func (s SingleFileSource) Backup() (backup_output.BackupOutput, error) {
	out, err := os.Open(s.Path)
	return backup_output.BackupOutput{Data: out}, err
}
