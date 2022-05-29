package mysql

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type MysqlSource struct {
	Host               string   `yaml:"Host"`
	Port               string   `yaml:"Port"`
	User               string   `yaml:"User"`
	Password           string   `yaml:"Password"`
	Database           string   `yaml:"Database"`
	MysqldumpExtraArgs []string `yaml:"MysqldumpExtraArgs"`
}

func (s MysqlSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MysqlSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MysqlSource need to have a Port")
	}
	if s.User == "" {
		return fmt.Errorf("MysqlSource need to have a User")
	}
	if s.Password == "" {
		return fmt.Errorf("MysqlSource need to have a Password")
	}
	if s.Database == "" {
		return fmt.Errorf("MysqlSource need to have a Database")
	}
	return nil
}

func (s MysqlSource) Backup() (backup_output.BackupOutput, error) {
	args := []string{
		"-h", s.Host,
		"-P", s.Port,
		"-u", s.User,
		"-p" + s.Password,
		s.Database,
	}
	args = append(s.MysqldumpExtraArgs, args...)
	return backup_process_utils.BackupProcessExecToFile(
		"mysqldump",
		args...,
	)
}
