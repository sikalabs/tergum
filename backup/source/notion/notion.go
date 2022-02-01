package notion

import (
	"fmt"
	"io"
	"os"

	"github.com/ondrejsika/notion-backup/lib/backup"
)

type NotionSource struct {
	Token   string `yaml:"Token"`
	SpaceID string `yaml:"SpaceID"`
	Format  string `yaml:"Format"`
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

func (s NotionSource) Backup() (io.ReadSeeker, error) {
	outputFile, err := os.CreateTemp("", "tergum-dump-notion-")
	outputFile.Seek(0, 0)
	backup.Backup(s.Token, s.SpaceID, s.Format, outputFile.Name())
	return outputFile, err
}
