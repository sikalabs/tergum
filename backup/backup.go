package backup

import (
	"errors"

	"github.com/sikalabs/tergum/backup_log"
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
	ID       string
	Name     string
	Mysql    mysql.Mysql
	FilePath filepath.FilePath
	File     file.File
	S3       s3.S3
}

type Backups struct {
	ID           string
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

func BackupAndSaveAll(backups []Backups) (backup_log.BackupGlobalLog, error) {
	var globalLog backup_log.BackupGlobalLog
	for i := 0; i < len(backups); i++ {
		backup := backups[i]
		data, err := Backup(backup.Source)
		if err != nil {
			globalLog.Logs = append(globalLog.Logs, backup_log.BackupLog{
				BackupID: backup.ID,
				Success:  false,
				Error:    err,
			})
		}
		for i := 0; i < len(backup.Destinations); i++ {
			destination := backup.Destinations[i]
			err := Save(destination, data)
			if err != nil {
				globalLog.Logs = append(globalLog.Logs, backup_log.BackupLog{
					BackupID:      backup.ID,
					DestinationID: destination.ID,
					Success:       false,
					Error:         err,
				})
			} else {
				globalLog.Logs = append(globalLog.Logs, backup_log.BackupLog{
					BackupID:      backup.ID,
					DestinationID: destination.ID,
					Success:       true,
					Error:         nil,
				})
			}
		}
	}
	backup_log.GlobalLogToOutput(globalLog)
	return globalLog, nil
}
