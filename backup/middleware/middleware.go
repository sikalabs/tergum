package middleware

import (
	"fmt"

	"github.com/sikalabs/tergum/backup/middleware/gzip"
)

type Middleware struct {
	Gzip *gzip.GzipMiddleware
}

func (m Middleware) Validate() error {
	if m.Gzip != nil {
		return m.Gzip.Validate()
	}

	return fmt.Errorf("no middleware detected")
}

func (m Middleware) Process(data []byte) ([]byte, error) {
	if m.Gzip != nil {
		return m.Gzip.Process(data)
	}

	return nil, fmt.Errorf("no middleware detected")
}
