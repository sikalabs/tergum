package postgres

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/sikalabs/tergum/utils/temp_utils"
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

	outputFileName := temp_utils.GetTempFileName()
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

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

	outputFileReader, err := os.Open(outputFileName)
	return outputFileReader, err
}
