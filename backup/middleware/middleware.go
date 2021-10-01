package middleware

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/middleware/gzip"
)

type Middleware struct {
	Gzip *gzip.GzipMiddleware `yaml:"Gzip"`
}

func (m Middleware) Validate() error {
	if m.Gzip != nil {
		return m.Gzip.Validate()
	}

	return fmt.Errorf("no middleware detected")
}

func (m Middleware) Process(data io.Reader) (io.Reader, error) {
	if m.Gzip != nil {
		return m.Gzip.Process(data)
	}

	return nil, fmt.Errorf("no middleware detected")
}
