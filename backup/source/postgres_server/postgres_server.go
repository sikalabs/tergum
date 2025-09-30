package postgres_server

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type PostgresServerSource struct {
	Host               string   `yaml:"Host" json:"Host,omitempty"`
	Port               string   `yaml:"Port" json:"Port,omitempty"`
	User               string   `yaml:"User" json:"User,omitempty"`
	Password           string   `yaml:"Password" json:"Password,omitempty"`
	PgdumpallExtraArgs []string `yaml:"PgdumpallExtraArgs" json:"PgdumpallExtraArgs,omitempty"`
	SSLMode            string   `yaml:"SSLMode" json:"SSLMode,omitempty"`
	UseBinaryBackup    bool     `yaml:"UseBinaryBackup" json:"UseBinaryBackup,omitempty"`
}

func (s PostgresServerSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("PostgresServerSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("PostgresServerSource requires Port")
	}
	if s.User == "" {
		return fmt.Errorf("PostgresServerSource requires User")
	}
	if s.Password == "" {
		return fmt.Errorf("PostgresServerSource requires Password")
	}
	return nil
}

func (s PostgresServerSource) Backup() (backup_output.BackupOutput, error) {
	env := os.Environ()
	env = append(env, "PGPASSWORD="+s.Password)
	if s.SSLMode != "" {
		env = append(env, "PGSSLMODE="+s.SSLMode)
	}

	var args []string
	var command string

	if s.UseBinaryBackup {
		command = "pg_basebackup"
		args = []string{
			"--host", s.Host,
			"--port", s.Port,
			"--user", s.User,
			"--no-password",
			"--format=tar",
			"--compress=9",
			"--pgdata=-",
		}
		args = append(s.PgdumpallExtraArgs, args...)
	} else {
		command = "pg_dumpall"
		args = []string{
			"--host", s.Host,
			"--port", s.Port,
			"--user", s.User,
			"--no-password",
		}
		args = append(s.PgdumpallExtraArgs, args...)
	}

	return backup_process_utils.BackupProcessExecEnvToFile(
		env,
		command,
		args...,
	)
}
