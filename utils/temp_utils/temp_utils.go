package temp_utils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sikalabs/tergum/utils/rand_utils"
)

func GetTempFileName() string {
	return filepath.Join(
		os.TempDir(),
		"tergum-tmp-"+rand_utils.GetRandString(10),
	)
}

func GetTempFilePipe() (io.Writer, io.Reader, error) {
	filename := GetTempFileName()
	w, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	defer w.Close()
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	return w, r, nil
}
