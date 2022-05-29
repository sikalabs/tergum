package ftp

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sikalabs/tergum/backup/backup_process"
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

	tmpDir, err := os.MkdirTemp("", "tergum-ftp-wget-")
	if err != nil {
		return nil, "", err
	}

	f, err := os.CreateTemp("", "tergum-")
	if err != nil {
		return nil, "", err
	}
	defer os.Remove(f.Name())

	bp := backup_process.BackupProcess{}
	bp.Init()

	// Download from FTP server using wget
	bp.ExecDirWait(
		tmpDir,
		"wget",
		"-q",
		"-r",
		"--user", s.User,
		"--password", s.Password,
		"ftp://"+s.Host,
	)
	if err != nil {
		return nil, "", err
	}

	// Create tar achrive
	bp.ExecDirWait(
		path.Join(tmpDir, s.Host),
		"tar",
		"-cf",
		f.Name(),
		".",
	)
	if err != nil {
		return nil, "", err
	}

	_, err = f.Seek(0, 0)
	return f, "", err
}
