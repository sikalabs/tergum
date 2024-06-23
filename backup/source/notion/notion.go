package notion

import (
	"fmt"
	"os"

	"github.com/ondrejsika/notion-backup/lib/backup"
	"github.com/sikalabs/tergum/backup_output"
)

type NotionSource struct {
	Token   string `yaml:"Token" json:"Token,omitempty"`
	SpaceID string `yaml:"SpaceID" json:"SpaceID,omitempty"`
	Format  string `yaml:"Format" json:"Format,omitempty"`
}

func (s NotionSource) Validate() error {
	if s.Token == "" {
		return fmt.Errorf("NotionSource need to have a Token")
	}
	if s.SpaceID == "" {
		return fmt.Errorf("NotionSource need to have a SpaceID")
	}
	if s.Format == "" {
		return fmt.Errorf("NotionSource need to have a Format")
	}
	if !(s.Format == "html" || s.Format == "markdown") {
		return fmt.Errorf("NotionSource.Format must be \"html\" or \"markdown\"")
	}
	return nil
}

func (s NotionSource) Backup() (backup_output.BackupOutput, error) {
	outputFile, err := os.CreateTemp("", "tergum-dump-notion-")
	outputFile.Seek(0, 0)
	backup.Backup(s.Token, s.SpaceID, s.Format, outputFile.Name())
	return backup_output.BackupOutput{Data: outputFile}, err
}
