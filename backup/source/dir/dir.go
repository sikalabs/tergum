package dir

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type DirSource struct {
	Path             string `yaml:"Path" json:"Path,omitempty"`
	IgnoreFailedRead bool   `yaml:"IgnoreFailedRead" json:"IgnoreFailedRead,omitempty"`
}

func (s DirSource) Validate() error {
	if s.Path == "" {
		return fmt.Errorf("DirSource need to have a Path")
	}
	return nil
}

func (s DirSource) Backup() (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	// Check if source path exists
	if _, err = os.Stat(s.Path); os.IsNotExist(err) {
		return bo, err
	}

	f, err := os.CreateTemp("", "tergum-")
	if err != nil {
		return bo, err
	}
	defer os.Remove(f.Name())

	args := []string{}
	if s.IgnoreFailedRead {
		args = append(args, "--ignore-failed-read")
	}
	args = append(args, []string{"-cf", f.Name(), s.Path}...)

	tmpBo, err := backup_process_utils.BackupProcessExecToFile(
		"tar",
		args...,
	)
	if err != nil {
		return tmpBo, err
	}

	// Seek to start of backup file
	_, err = f.Seek(0, 0)
	if err != nil {
		return bo, err
	}

	bo = backup_output.BackupOutput{
		Data:   f,
		Stderr: tmpBo.Stderr,
	}

	return bo, nil
}
