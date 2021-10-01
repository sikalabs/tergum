package mysql_server

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/sikalabs/tergum/utils/temp_utils"
)

type MysqlServerSource struct {
	Host               string   `yaml:"Host"`
	Port               string   `yaml:"Port"`
	User               string   `yaml:"User"`
	Password           string   `yaml:"Password"`
	MysqldumpExtraArgs []string `yaml:"MysqldumpExtraArgs"`
}

func (s MysqlServerSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MysqlServerSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MysqlServerSource need to have a Port")
	}
	if s.User == "" {
		return fmt.Errorf("MysqlServerSource need to have a User")
	}
	if s.Password == "" {
		return fmt.Errorf("MysqlServerSource need to have a Password")
	}
	return nil
}

func (s MysqlServerSource) Backup() (io.ReadSeeker, error) {
	var err error

	outputFileName := temp_utils.GetTempFileName()
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	args := []string{
		"-h", s.Host,
		"-P", s.Port,
		"-u", s.User,
		"-p" + s.Password,
		"--all-databases",
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
	cmd.Wait()

	outputFileReader, err := os.Open(outputFileName)
	return outputFileReader, err
}
