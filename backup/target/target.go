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
)

type Target struct {
	ID          string                      `yaml:"ID"`
	Middlewares []middleware.Middleware     `yaml:"Middlewares"`
	S3          *s3.S3Target                `yaml:"S3"`
	File        *file.FileTarget            `yaml:"File"`
	FilePath    *filepath.FilePathTarget    `yaml:"FilePath"`
	AzureBlob   *azure_blob.AzureBlobTarget `yaml:"AzureBlob"`
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

	return ""
}
