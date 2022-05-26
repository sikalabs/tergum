package ftp

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

type FTPSource struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

func (s FTPSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("FTPSource need to have a Host")
	}
	if s.User == "" {
		return fmt.Errorf("FTPSource requires User")
	}
	if s.Password == "" {
		return fmt.Errorf("FTPSource requires Password")
	}
	return nil
}

func (s FTPSource) Backup() (io.ReadSeeker, string, error) {
	var err error

	errorMessage := new(strings.Builder)

	wgetDir, err := os.MkdirTemp("", "tergum-ftp-wget-")
	if err != nil {
		return nil, "", err
	}

	cmd := exec.Command(
		"wget",
		"-q",
		"-r",
		"--user", s.User,
		"--password", s.Password,
		"ftp://"+s.Host,
	)
	cmd.Dir = wgetDir
	cmd.Stderr = errorMessage

	err = cmd.Start()
	if err != nil {
		return nil, errorMessage.String(), err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, errorMessage.String(), err
	}

	outputFile, err := os.CreateTemp("", "tergum-tar-gz-")
	if err != nil {
		return nil, "", err
	}
	defer os.Remove(outputFile.Name())

	wgetDataRoot := path.Join(wgetDir, s.Host)

	cmd = exec.Command(
		"tar",
		"-cf",
		outputFile.Name(),
		".",
	)
	cmd.Dir = wgetDataRoot

	err = cmd.Start()
	if err != nil {
		return nil, errorMessage.String(), err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, errorMessage.String(), err
	}

	outputFile.Seek(0, 0)
	return outputFile, errorMessage.String(), err
}
