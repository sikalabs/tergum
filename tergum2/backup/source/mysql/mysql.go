package mysql

import (
	"fmt"
	"os/exec"
)

type MysqlSource struct {
	Host               string
	Port               string
	User               string
	Password           string
	Database           string
	MysqldumpExtraArgs []string
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

func (s MysqlSource) Backup() ([]byte, error) {
	args := []string{
		"-h", s.Host,
		"-P", s.Port,
		"-u", s.User,
		"-p" + s.Password,
		s.Database,
	}
	cmd := exec.Command(
		"mysqldump",
		append(s.MysqldumpExtraArgs, args...)...,
	)
	out, err := cmd.Output()
	return out, err
}
