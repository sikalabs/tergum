package dir

import (
	"fmt"
	"io"
	"os"

	"github.com/sikalabs/tergum/backup_process_utils"
)

type DirSource struct {
	Path string `yaml:"Path"`
}

func (s DirSource) Validate() error {
	if s.Path == "" {
		return fmt.Errorf("DirSource need to have a Path")
	}
	return nil
}

func (s DirSource) Backup() (io.ReadSeeker, string, error) {
	var err error

	// Check if source path exists
	if _, err = os.Stat(s.Path); os.IsNotExist(err) {
		return nil, "", err
	}

	f, err := os.CreateTemp("", "tergum-")
	if err != nil {
		return nil, "", err
	}
	defer os.Remove(f.Name())

	_, stderr, err := backup_process_utils.BackupProcessExecToFile(
		"tar",
		"-cf",
		f.Name(),
		s.Path,
	)
	if err != nil {
		return nil, "", err
	}

	// Seek to start of backup file
	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, "", err
	}

	return f, stderr, nil
}
