package postgres_server

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type PostgresServerSource struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

func (s PostgresServerSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("PostgresServerSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("PostgresServerSource requires Port")
	}
	if s.User == "" {
		return fmt.Errorf("PostgresServerSource requires User")
	}
	if s.Password == "" {
		return fmt.Errorf("PostgresServerSource requires Password")
	}
	return nil
}

func (s PostgresServerSource) Backup() (io.ReadSeeker, string, error) {
	var err error
	errorMessage := new(strings.Builder)

	outputFile, err := os.CreateTemp("", "tergum-dump-postgres-")
	if err != nil {
		return nil, "", err
	}
	defer os.Remove(outputFile.Name())

	cmd := exec.Command(
		"pg_dumpall",
		"--host", s.Host,
		"--port", s.Port,
		"--user", s.User,
		"--no-password",
	)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PGPASSWORD="+s.Password)
	cmd.Stdout = outputFile
	cmd.Stderr = errorMessage

	err = cmd.Start()
	if err != nil {
		return nil, errorMessage.String(), err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return nil, errorMessage.String(), err
	}

	outputFile.Seek(0, 0)
	return outputFile, "", nil
}
