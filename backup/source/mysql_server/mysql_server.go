package mysql_server

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/backup_process_utils"
)

type MysqlServerSource struct {
	Host               string   `yaml:"Host"`
	Port               string   `yaml:"Port"`
	User               string   `yaml:"User"`
	Password           string   `yaml:"Password"`
	MysqldumpExtraArgs []string `yaml:"MysqldumpExtraArgs"`
}

func (s MysqlServerSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MysqlServerSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MysqlServerSource need to have a Port")
	}
	if s.User == "" {
		return fmt.Errorf("MysqlServerSource need to have a User")
	}
	if s.Password == "" {
		return fmt.Errorf("MysqlServerSource need to have a Password")
	}
	return nil
}

func (s MysqlServerSource) Backup() (io.ReadSeeker, string, error) {
	args := []string{
		"-h", s.Host,
		"-P", s.Port,
		"-u", s.User,
		"-p" + s.Password,
		"--all-databases",
	}
	args = append(s.MysqldumpExtraArgs, args...)
	return backup_process_utils.BackupProcessExecToFile(
		"mysqldump",
		args...,
	)
}
