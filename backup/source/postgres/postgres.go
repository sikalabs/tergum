package postgres

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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

func (s PostgresSource) Backup() (io.ReadSeeker, error) {
	var err error

	outputFile, err := os.CreateTemp("", "tergum-dump-postgres-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(outputFile.Name())

	cmd := exec.Command(
		"pg_dump",
		"host="+s.Host+
			" port="+s.Port+
			" user="+s.User+
			" password="+s.Password+
			" dbname="+s.Database,
	)
	cmd.Stdout = outputFile

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	cmd.Wait()

	outputFile.Seek(0, 0)
	return outputFile, nil
}
