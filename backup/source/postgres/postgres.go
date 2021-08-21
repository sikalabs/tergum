package postgres

import (
	"fmt"
	"os/exec"
)

type PostgresSource struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
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

func (s PostgresSource) Backup() ([]byte, error) {
	cmd := exec.Command(
		"pg_dump",
		"host="+s.Host+
			" port="+s.Port+
			" user="+s.User+
			" password="+s.Password+
			" dbname="+s.Database,
	)
	out, err := cmd.Output()
	return out, err
}
