package postgres

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type PostgresSource struct {
	Host            string   `yaml:"Host" json:"Host,omitempty"`
	Port            string   `yaml:"Port" json:"Port,omitempty"`
	User            string   `yaml:"User" json:"User,omitempty"`
	Password        string   `yaml:"Password" json:"Password,omitempty"`
	Database        string   `yaml:"Database" json:"Database,omitempty"`
	PgdumpExtraArgs []string `yaml:"PgdumpExtraArgs" json:"PgdumpExtraArgs,omitempty"`
	SSLMode         string   `yaml:"SSLMode" json:"SSLMode,omitempty"`
}

func (s PostgresSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("PostgresSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("PostgresSource requires Port")
	}
	if s.User == "" {
		return fmt.Errorf("PostgresSource requires User")
	}
	if s.Password == "" {
		return fmt.Errorf("PostgresSource requires Password")
	}
	if s.Database == "" {
		return fmt.Errorf("PostgresSource requires Database")
	}
	return nil
}

func (s PostgresSource) Backup() (backup_output.BackupOutput, error) {
	args := []string{
		"host=" + s.Host +
			" port=" + s.Port +
			" user=" + s.User +
			" password=" + s.Password +
			" dbname=" + s.Database,
	}
	args = append(s.PgdumpExtraArgs, args...)

	env := os.Environ()
	if s.SSLMode != "" {
		env = append(env, "PGSSLMODE="+s.SSLMode)
	}

	return backup_process_utils.BackupProcessExecEnvToFile(
		env,
		"pg_dump",
		args...,
	)
}
