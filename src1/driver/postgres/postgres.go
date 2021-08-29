package postgres

import (
	"errors"
	"os/exec"
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func ValidatePostgres(config Postgres) error {
	if config.Host == "" {
		return errors.New("postgres requires host")
	}
	if config.Port == "" {
		return errors.New("postgres requires port")
	}
	if config.User == "" {
		return errors.New("postgres requires user")
	}
	if config.Password == "" {
		return errors.New("postgres requires password")
	}
	if config.Database == "" {
		return errors.New("postgres requires database")
	}
	return nil
}

func BackupPostgres(config Postgres) ([]byte, error) {
	cmd := exec.Command(
		"pg_dump",
		"host="+config.Host+
			" port="+config.Port+
			" user="+config.User+
			" password="+config.Password+
			" dbname="+config.Database,
	)
	out, err := cmd.Output()
	return out, err
}
