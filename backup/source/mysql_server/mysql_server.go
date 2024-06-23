package mysql_server

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type MysqlServerSource struct {
	Host               string   `yaml:"Host" json:"Host,omitempty"`
	Port               string   `yaml:"Port" json:"Port,omitempty"`
	User               string   `yaml:"User" json:"User,omitempty"`
	Password           string   `yaml:"Password" json:"Password,omitempty"`
	MysqldumpExtraArgs []string `yaml:"MysqldumpExtraArgs" json:"MysqldumpExtraArgs,omitempty"`
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

func (s MysqlServerSource) Backup() (backup_output.BackupOutput, error) {
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
