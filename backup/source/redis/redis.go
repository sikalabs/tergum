package redis

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type RedisSource struct {
	Host string `yaml:"Host" json:"Host,omitempty"`
	Port string `yaml:"Port" json:"Port,omitempty"`
}

func (s RedisSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("RedisSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("RedisSource need to have a Port")
	}
	return nil
}

func (s RedisSource) Backup() (backup_output.BackupOutput, error) {
	args := []string{
		"-h", s.Host,
		"-p", s.Port,
		"--rdb", "-",
	}
	return backup_process_utils.BackupProcessExecToFile(
		"redis-cli",
		args...,
	)
}
