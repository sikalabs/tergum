package gzip

import (
	"io"

	"github.com/sikalabs/tergum/utils/gzip_utils"
)

type GzipMiddleware struct{}

func (m GzipMiddleware) Validate() error {
	return nil
}

func (m GzipMiddleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	return gzip_utils.GzipIO(data)
}
