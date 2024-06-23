package dummy

import (
	"bytes"
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
)

type DummySource struct {
	Content string `yaml:"Content" json:"Content,omitempty"`
}

func (s DummySource) Validate() error {
	if s.Content == "" {
		return fmt.Errorf("DummySource need to have a Content")
	}
	return nil
}

func (s DummySource) Backup() (backup_output.BackupOutput, error) {
	out := bytes.NewReader([]byte(s.Content))
	return backup_output.BackupOutput{Data: out}, nil
}
