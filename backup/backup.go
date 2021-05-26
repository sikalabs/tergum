package backup

import (
	"errors"
	"log"

	"github.com/sikalabs/tergum/driver/file"
	"github.com/sikalabs/tergum/driver/filepath"
	"github.com/sikalabs/tergum/driver/mysql"
	"github.com/sikalabs/tergum/driver/s3"
)

type BackupSource struct {
	Name  string
	Mysql mysql.Mysql
}

type BackupDestination struct {
	Name     string
	Mysql    mysql.Mysql
	FilePath filepath.FilePath
	File     file.File
	S3       s3.S3
}

type Backups struct {
	Source       BackupSource
	Destinations []BackupDestination
}

func Backup(config BackupSource) ([]byte, error) {
	switch config.Name {
	case "mysql":
		err := mysql.ValidateMysql(config.Mysql)
		if err != nil {
			return nil, err
		}
		return mysql.BackupMysql(config.Mysql)
	}
	return nil, errors.New("no backup driver found")
}

func Save(config BackupDestination, data []byte) error {
	switch config.Name {
	case "filepath":
		err := filepath.ValidateFilePath(config.FilePath)
		if err != nil {
			return err
		}
		return filepath.SaveFilePath(config.FilePath, data)
	case "file":
		err := file.ValidateFile(config.File)
		if err != nil {
			return err
		}
		return file.SaveFile(config.File, data)
	case "s3":
		err := s3.ValidateS3(config.S3)
		if err != nil {
			return err
		}
		return s3.SaveS3(config.S3, data)
	}
	return errors.New("no backup target driver found")
}

func BackupAndSaveAll(backups []Backups) error {
	for i := 0; i < len(backups); i++ {
		backup := backups[i]
		data, err := Backup(backup.Source)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(backup.Destinations); i++ {
			destination := backup.Destinations[i]
			err := Save(destination, data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
