package mssql

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process"
)

type MssqlSource struct {
	Host                   string   `yaml:"Host" json:"Host,omitempty"`
	Port                   string   `yaml:"Port" json:"Port,omitempty"`
	User                   string   `yaml:"User" json:"User,omitempty"`
	Password               string   `yaml:"Password" json:"Password,omitempty"`
	Database               string   `yaml:"Database" json:"Database,omitempty"`
	TrustServerCertificate bool     `yaml:"TrustServerCertificate" json:"TrustServerCertificate,omitempty"`
	ExtractAllTableData    bool     `yaml:"ExtractAllTableData" json:"ExtractAllTableData,omitempty"`
	SqlpackageExtraArgs    []string `yaml:"SqlpackageExtraArgs" json:"SqlpackageExtraArgs,omitempty"`
}

func (s MssqlSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MssqlSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MssqlSource need to have a Port")
	}
	if s.User == "" {
		return fmt.Errorf("MssqlSource need to have a User")
	}
	if s.Password == "" {
		return fmt.Errorf("MssqlSource need to have a Password")
	}
	if s.Database == "" {
		return fmt.Errorf("MssqlSource need to have a Database")
	}
	return nil
}

func (s MssqlSource) Backup() (backup_output.BackupOutput, error) {
	var bo backup_output.BackupOutput

	bp := backup_process.BackupProcess{}
	bp.Init()

	tempDir, err := os.MkdirTemp("", "tergum-mssql-*")
	if err != nil {
		return bo, err
	}
	defer os.RemoveAll(tempDir)

	targetFile := tempDir + "/" + s.Database + ".dacpac"

	args := []string{
		"/Action:Extract",
		"/SourceServerName:" + s.Host + "," + s.Port,
		"/SourceDatabaseName:" + s.Database,
		"/SourceUser:" + s.User,
		"/SourcePassword:" + s.Password,
		"/TargetFile:" + targetFile,
	}
	if s.TrustServerCertificate {
		args = append(args, "/SourceTrustServerCertificate:True")
	}
	if s.ExtractAllTableData {
		args = append(args, "/p:ExtractAllTableData=True")
	}
	args = append(args, s.SqlpackageExtraArgs...)

	err = bp.ExecWait("sqlpackage", args...)
	stderr, _ := bp.GetStderr()
	if err != nil {
		return backup_output.BackupOutput{Stderr: stderr}, err
	}

	f, err := os.Open(targetFile)
	if err != nil {
		return backup_output.BackupOutput{Stderr: stderr}, err
	}

	return backup_output.BackupOutput{Data: f, Stderr: stderr}, nil
}
