package middleware

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/middleware/gzip"
	"github.com/sikalabs/tergum/backup/middleware/symmetric_encryption"
)

type Middleware struct {
	Gzip                *gzip.GzipMiddleware                                `yaml:"Gzip"`
	SymmetricEncryption *symmetric_encryption.SymmetricEncryptionMiddleware `yaml:"SymmetricEncryption"`
}

func (m Middleware) Validate() error {
	if m.Gzip != nil {
		return m.Gzip.Validate()
	}

	if m.SymmetricEncryption != nil {
		return m.SymmetricEncryption.Validate()
	}

	return fmt.Errorf("no middleware detected")
}

func (m Middleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	if m.Gzip != nil {
		return m.Gzip.Process(data)
	}

	if m.SymmetricEncryption != nil {
		return m.SymmetricEncryption.Process(data)
	}

	return nil, fmt.Errorf("no middleware detected")
}

func (m Middleware) Name() string {
	if m.Gzip != nil {
		return "Gzip"
	}

	if m.SymmetricEncryption != nil {
		return "SymmetricEncryption"
	}

	return ""
}
