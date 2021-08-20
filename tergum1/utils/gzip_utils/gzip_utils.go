package gzip_utils

import (
	"bytes"
	"compress/gzip"
	"os"
)

func WriteGzipFile(path string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		return err
	}

	w := gzip.NewWriter(f)

	_, err = w.Write(data)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func GzipBytes(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	err = gz.Close()
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
