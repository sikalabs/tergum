package mongo

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type MongoSource struct {
	Host                   string `yaml:"Host"`
	Port                   string `yaml:"Port"`
	User                   string `yaml:"User"`
	Password               string `yaml:"Password"`
	Database               string `yaml:"Database"`
	AuthenticationDatabase string `yaml:"AuthenticationDatabase"`
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

func (s MongoSource) Backup() (io.ReadSeeker, string, error) {
	outputFile, err := os.CreateTemp("", "tergum-dump-mongo-")
	errorMessage := new(strings.Builder)

	if err != nil {
		return nil, "", err
	}
	defer os.Remove(outputFile.Name())
	args := []string{
		"--archive=" + outputFile.Name(),
		"--host", s.Host,
		"--port", s.Port,
	}
	if s.User != "" {
		// Default AuthenticationDatabase is admin
		if s.AuthenticationDatabase == "" {
			s.AuthenticationDatabase = "admin"
		}
		args = append(
			args,
			"--username", s.User,
			"--password", s.Password,
			"--authenticationDatabase", s.AuthenticationDatabase,
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
	cmd.Stderr = errorMessage

	_, err = cmd.Output()
	outputFile.Seek(0, 0)
	return outputFile, errorMessage.String(), err
}
