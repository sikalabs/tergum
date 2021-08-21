package backup

import (
	"errors"

	"github.com/sikalabs/tergum/tergum1/backup_log"
	"github.com/sikalabs/tergum/tergum1/driver/file"
	"github.com/sikalabs/tergum/tergum1/driver/filepath"
	"github.com/sikalabs/tergum/tergum1/driver/mysql"
	"github.com/sikalabs/tergum/tergum1/driver/postgres"
	"github.com/sikalabs/tergum/tergum1/driver/s3"
	"github.com/sikalabs/tergum/tergum1/middleware"
	"github.com/sikalabs/tergum/tergum1/utils/gzip_utils"
)

type BackupSource struct {
	Name     string
	Mysql    mysql.Mysql
	Postgres postgres.Postgres
}

type BackupDestination struct {
	ID          string
	Name        string
	Middlewares []middleware.Middleware
	Mysql       mysql.Mysql
	FilePath    filepath.FilePath
	File        file.File
	S3          s3.S3
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
	case "postgres":
		err := postgres.ValidatePostgres(config.Postgres)
		if err != nil {
			return nil, err
		}
		return postgres.BackupPostgres(config.Postgres)
	}
	return nil, errors.New("no backup driver found")
}

func Transform(middleware middleware.Middleware, data []byte) ([]byte, error) {
	switch middleware.Name {
	case "gzip":
		data, err := gzip_utils.GzipBytes(data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("no middleware found")
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
			continue
		}
		for i := 0; i < len(backup.Destinations); i++ {
			destination := backup.Destinations[i]
			transformOK := true
			for j := 0; j < len(destination.Middlewares); j++ {
				mw := destination.Middlewares[j]
				data, err = Transform(mw, data)
				if err != nil {
					transformOK = false
					globalLog.Logs = append(globalLog.Logs, backup_log.BackupLog{
						BackupID:      backup.ID,
						DestinationID: destination.ID,
						Success:       false,
						Error:         err,
					})
				}
			}
			if !transformOK {
				continue
			}
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
