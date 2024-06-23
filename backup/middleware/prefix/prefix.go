package prefix

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type PrefixMiddleware struct {
	Prefix string `yaml:"Prefix" json:"Prefix,omitempty"`
}

func (m PrefixMiddleware) Validate() error {
	if m.Prefix == "" {
		return fmt.Errorf("Prefix must be defined")
	}
	return nil
}

func (m PrefixMiddleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	var err error
	f, _ := os.CreateTemp("", "tergum-")
	_, err = io.Copy(f, strings.NewReader(m.Prefix))
	_ = err
	_, err = io.Copy(f, data)
	f.Seek(0, 0)
	return f, err
}
