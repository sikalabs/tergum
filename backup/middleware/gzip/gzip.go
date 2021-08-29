package gzip

import "github.com/sikalabs/tergum/utils/gzip_utils"

type GzipMiddleware struct{}

func (m GzipMiddleware) Validate() error {
	return nil
}

func (m GzipMiddleware) Process(data []byte) ([]byte, error) {
	return gzip_utils.GzipBytes(data)
}
