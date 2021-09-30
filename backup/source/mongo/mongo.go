package mongo

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/sikalabs/tergum/utils/temp_utils"
)

type MongoSource struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
}

func (s MongoSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MongoSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MongoSource need to have a Port")
	}
	return nil
}

func (s MongoSource) Backup() ([]byte, error) {
	outputFile := temp_utils.GetTempFileName()
	args := []string{
		"--archive=" + outputFile,
		"--host", s.Host,
		"--port", s.Port,
	}
	if s.User != "" {
		args = append(
			args,
			"--username", s.User,
			"--password", s.Password,
		)
	}
	if s.Database != "" {
		args = append(
			args,
			"--db", s.Database,
		)
	}
	cmd := exec.Command(
		"mongodump",
		args...,
	)
	cmd.Output()
	out, err := ioutil.ReadFile(outputFile)
	return out, err
}
