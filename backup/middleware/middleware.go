package middleware

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/middleware/gzip"
	"github.com/sikalabs/tergum/backup/middleware/prefix"
	"github.com/sikalabs/tergum/backup/middleware/suffix"
	"github.com/sikalabs/tergum/backup/middleware/symmetric_encryption"
)

type Middleware struct {
	Gzip                *gzip.GzipMiddleware                                `yaml:"Gzip"`
	SymmetricEncryption *symmetric_encryption.SymmetricEncryptionMiddleware `yaml:"SymmetricEncryption"`
	Prefix              *prefix.PrefixMiddleware                            `yaml:"Prefix"`
	Suffix              *suffix.SuffixMiddleware                            `yaml:"Suffix"`
}

func (m Middleware) Validate() error {
	if m.Gzip != nil {
		return m.Gzip.Validate()
	}

	if m.SymmetricEncryption != nil {
		return m.SymmetricEncryption.Validate()
	}

	if m.Prefix != nil {
		return m.Prefix.Validate()
	}

	if m.Suffix != nil {
		return m.Suffix.Validate()
	}

	return fmt.Errorf("validate: no middleware detected")
}

func (m Middleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	if m.Gzip != nil {
		return m.Gzip.Process(data)
	}

	if m.SymmetricEncryption != nil {
		return m.SymmetricEncryption.Process(data)
	}

	if m.Prefix != nil {
		return m.Prefix.Process(data)
	}

	if m.Suffix != nil {
		return m.Suffix.Process(data)
	}

	return nil, fmt.Errorf("process: no middleware detected")
}

func (m Middleware) Name() string {
	if m.Gzip != nil {
		return "Gzip"
	}

	if m.SymmetricEncryption != nil {
		return "SymmetricEncryption"
	}

	if m.Prefix != nil {
		return "Prefix"
	}

	if m.Suffix != nil {
		return "Suffix"
	}

	return ""
}
