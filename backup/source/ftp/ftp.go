package ftp

import (
	"fmt"
	"os"
	"path"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process"
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

func (s FTPSource) Backup() (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	tmpDir, err := os.MkdirTemp("", "tergum-ftp-wget-")
	if err != nil {
		return bo, err
	}

	f, err := os.CreateTemp("", "tergum-")
	if err != nil {
		return bo, err
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
		return bo, err
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
		return bo, err
	}

	_, err = f.Seek(0, 0)

	bo = backup_output.BackupOutput{
		Data: f,
	}

	return bo, err
}
