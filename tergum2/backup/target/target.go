package target

import (
	"fmt"

	"github.com/sikalabs/tergum/tergum2/backup/middleware"
	"github.com/sikalabs/tergum/tergum2/backup/target/file"
	"github.com/sikalabs/tergum/tergum2/backup/target/filepath"
	"github.com/sikalabs/tergum/tergum2/backup/target/s3"
)

type Target struct {
	Middlewares []middleware.Middleware
	S3          *s3.S3Target
	File        *file.FileTarget
	FilePath    *filepath.FilePathTarget
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

	return fmt.Errorf("no target detected")
}

func (t Target) Save(data []byte) error {
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

	return fmt.Errorf("no target detected")
}
