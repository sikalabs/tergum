package mysql

import (
	"errors"
	"os/exec"
)

type Mysql struct {
	Host               string
	Port               string
	User               string
	Password           string
	Database           string
	MysqldumpExtraArgs []string
}

func ValidateMysql(config Mysql) error {
	if config.Host == "" {
		return errors.New("mysql requires host")
	}
	if config.Port == "" {
		return errors.New("mysql requires port")
	}
	if config.User == "" {
		return errors.New("mysql requires user")
	}
	if config.Password == "" {
		return errors.New("mysql requires password")
	}
	if config.Database == "" {
		return errors.New("mysql requires database")
	}
	return nil
}

func BackupMysql(config Mysql) ([]byte, error) {
	args := []string{
		"-h", config.Host,
		"-P", config.Port,
		"-u", config.User,
		"-p" + config.Password,
		config.Database,
	}
	cmd := exec.Command(
		"mysqldump",
		append(config.MysqldumpExtraArgs, args...)...,
	)
	out, err := cmd.Output()
	return out, err
}
