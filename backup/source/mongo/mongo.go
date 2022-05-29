package mongo

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type MongoSource struct {
	Host                   string `yaml:"Host"`
	Port                   string `yaml:"Port"`
	User                   string `yaml:"User"`
	Password               string `yaml:"Password"`
	Database               string `yaml:"Database"`
	AuthenticationDatabase string `yaml:"AuthenticationDatabase"`
}

func (s MongoSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MongoSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MongoSource need to have a Port")
	}
	return nil
}

func (s MongoSource) Backup() (backup_output.BackupOutput, error) {
	var bo backup_output.BackupOutput

	// Create file for backup
	f, err := os.CreateTemp("", "tergum-")
	if err != nil {
		return bo, err
	}
	defer os.Remove(f.Name())

	args := []string{
		"--archive=" + f.Name(),
		"--host", s.Host,
		"--port", s.Port,
	}
	if s.User != "" {
		// Default AuthenticationDatabase is admin
		if s.AuthenticationDatabase == "" {
			s.AuthenticationDatabase = "admin"
		}
		args = append(
			args,
			"--username", s.User,
			"--password", s.Password,
			"--authenticationDatabase", s.AuthenticationDatabase,
		)
	}
	if s.Database != "" {
		args = append(
			args,
			"--db", s.Database,
		)
	}

	tmpBo, err := backup_process_utils.BackupProcessExecToFile(
		"mongodump",
		args...,
	)
	if err != nil {
		return bo, err
	}

	// Seek to start of backup file
	_, err = f.Seek(0, 0)
	if err != nil {
		return bo, err
	}

	bo = backup_output.BackupOutput{
		Data:   f,
		Stderr: tmpBo.Stderr,
	}

	return bo, err
}
