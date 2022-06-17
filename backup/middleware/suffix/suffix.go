package suffix

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type SuffixMiddleware struct {
	Suffix string `yaml:"Suffix"`
}

func (m SuffixMiddleware) Validate() error {
	if m.Suffix == "" {
		return fmt.Errorf("Suffix must be defined")
	}
	return nil
}

func (m SuffixMiddleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	var err error
	f, _ := os.CreateTemp("", "tergum-")
	_, err = io.Copy(f, data)
	_ = err
	_, err = io.Copy(f, strings.NewReader(m.Suffix))
	f.Seek(0, 0)
	return f, err
}
