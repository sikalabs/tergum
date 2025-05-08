package target

import (
	"fmt"
	"io"
	"os"

	"github.com/sikalabs/tergum/backup/middleware"
	"github.com/sikalabs/tergum/backup/target/azure_blob"
	"github.com/sikalabs/tergum/backup/target/file"
	"github.com/sikalabs/tergum/backup/target/filepath"
	"github.com/sikalabs/tergum/backup/target/s3"
	"github.com/sikalabs/tergum/backup/target/telegram"
)

type Target struct {
	ID             string                      `yaml:"ID" json:"ID,omitempty"`
	Middlewares    []middleware.Middleware     `yaml:"Middlewares" json:"Middlewares,omitempty"`
	S3             *s3.S3Target                `yaml:"S3" json:"S3,omitempty"`
	File           *file.FileTarget            `yaml:"File" json:"File,omitempty"`
	FilePath       *filepath.FilePathTarget    `yaml:"FilePath" json:"FilePath,omitempty"`
	AzureBlob      *azure_blob.AzureBlobTarget `yaml:"AzureBlob" json:"AzureBlob,omitempty"`
	TelegramTarget *telegram.TelegramTarget    `yaml:"Telegram" json:"Telegram,omitempty"`
}

func (t Target) Validate() error {
	if t.S3 != nil {
		return t.S3.Validate()
	}

	if t.File != nil {
		return t.File.Validate()
	}

	if t.FilePath != nil {
		return t.FilePath.Validate()
	}

	if t.AzureBlob != nil {
		return t.AzureBlob.Validate()
	}

	if t.TelegramTarget != nil {
		return t.TelegramTarget.Validate()
	}

	return fmt.Errorf("target/validate: no target detected")
}

func (t Target) Save(data io.ReadSeeker) (int64, error) {
	size, _ := data.Seek(0, os.SEEK_END)
	data.Seek(0, 0)

	if t.S3 != nil {
		s3 := *t.S3
		return size, s3.Save(data)
	}

	if t.File != nil {
		f := *t.File
		return size, f.Save(data)
	}

	if t.FilePath != nil {
		fp := *t.FilePath
		return size, fp.Save(data)
	}

	if t.AzureBlob != nil {
		fp := *t.AzureBlob
		return size, fp.Save(data)
	}

	if t.TelegramTarget != nil {
		tg := *t.TelegramTarget
		return size, tg.Save(data)
	}

	return size, fmt.Errorf("target/save: no target detected")
}

func (t Target) Name() string {
	if t.S3 != nil {
		return "S3"
	}

	if t.File != nil {
		return "File"
	}

	if t.FilePath != nil {
		return "FilePath"
	}

	if t.AzureBlob != nil {
		return "AzureBlob"
	}

	if t.TelegramTarget != nil {
		return "Telegram"
	}

	return ""
}
