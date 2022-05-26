package dir

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type DirSource struct {
	Path string `yaml:"Path"`
}

func (s DirSource) Validate() error {
	if s.Path == "" {
		return fmt.Errorf("DirSource need to have a FilePath")
	}
	return nil
}

func (s DirSource) Backup() (io.ReadSeeker, string, error) {
	var err error

	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		return nil, "", err
	}

	outputFile, err := os.CreateTemp("", "tergum-tar-gz-")
	if err != nil {
		return nil, "", err
	}
	defer os.Remove(outputFile.Name())

	cmd := exec.Command(
		"tar",
		"-cf",
		outputFile.Name(),
		s.Path,
	)

	err = cmd.Start()
	if err != nil {
		return nil, "", err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, "", err
	}

	outputFile.Seek(0, 0)
	return outputFile, "", err
}
