package mysql

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type MysqlSource struct {
	Host               string   `yaml:"Host"`
	Port               string   `yaml:"Port"`
	User               string   `yaml:"User"`
	Password           string   `yaml:"Password"`
	Database           string   `yaml:"Database"`
	MysqldumpExtraArgs []string `yaml:"MysqldumpExtraArgs"`
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

func (s MysqlSource) Backup() (io.ReadSeeker, error) {
	var err error

	outputFile, err := os.CreateTemp("", "tergum-dump-mysql-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(outputFile.Name())

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
	cmd.Stdout = outputFile

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	outputFile.Seek(0, 0)
	return outputFile, nil
}
