package postgres

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/backup_process_utils"
)

type PostgresSource struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
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

func (s PostgresSource) Backup() (io.ReadSeeker, string, error) {
	return backup_process_utils.BackupProcessExecToFile(
		"pg_dump",
		"host="+s.Host+
			" port="+s.Port+
			" user="+s.User+
			" password="+s.Password+
			" dbname="+s.Database,
	)
}
