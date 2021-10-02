package gzip

import (
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/sikalabs/tergum/utils/gzip_utils"
)

type GzipMiddleware struct{}

func (m GzipMiddleware) Validate() error {
	return nil
}

func (m GzipMiddleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	var err error

	size, _ := data.Seek(0, os.SEEK_END)
	data.Seek(0, 0)

	bar := pb.Full.Start64(size)

	// create proxy reader
	barReader := bar.NewProxyReader(data)
	out, err := gzip_utils.GzipIO(barReader)
	bar.Finish()

	return out, err
}
