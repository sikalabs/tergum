package target

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/middleware"
	"github.com/sikalabs/tergum/backup/target/file"
	"github.com/sikalabs/tergum/backup/target/filepath"
	"github.com/sikalabs/tergum/backup/target/s3"
)

type Target struct {
	ID          string                   `yaml:"ID"`
	Middlewares []middleware.Middleware  `yaml:"Middlewares"`
	S3          *s3.S3Target             `yaml:"S3"`
	File        *file.FileTarget         `yaml:"File"`
	FilePath    *filepath.FilePathTarget `yaml:"FilePath"`
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

	return fmt.Errorf("target/validate: no target detected")
}

func (t Target) Save(data io.ReadSeeker) error {
	if t.S3 != nil {
		s3 := *t.S3
		return s3.Save(data)
	}

	if t.File != nil {
		f := *t.File
		return f.Save(data)
	}

	if t.FilePath != nil {
		fp := *t.FilePath
		return fp.Save(data)
	}

	return fmt.Errorf("target/save: no target detected")
}
