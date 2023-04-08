package gitlab

import (
	"fmt"
	"os"
	"time"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process"
	"github.com/sikalabs/tergum/utils/rand_utils"
)

type GitlabSource struct {
	NamePrefix string `yaml:"NamePrefix"`
}

func (s GitlabSource) Validate() error {
	if s.NamePrefix == "" {
		return fmt.Errorf("GitlabSource need to have a NamePrefix")
	}
	return nil
}

func (s GitlabSource) Backup() (backup_output.BackupOutput, error) {
	var bo backup_output.BackupOutput
	var err error

	backupName := fmt.Sprintf("%s_%d_%s", s.NamePrefix, time.Now().Unix(), rand_utils.GetRandString(3))
	backupFilePath := fmt.Sprintf("/var/opt/gitlab/backups/%s_gitlab_backup.tar", backupName)

	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.ExecWait("gitlab-backup", "create", "BACKUP="+backupName)
	stderr, _ := bp.GetStderr()
	if err != nil {
		return backup_output.BackupOutput{
			Data:   nil,
			Stderr: stderr,
		}, err
	}

	f, err := os.Open(backupFilePath)
	if err != nil {
		return backup_output.BackupOutput{
			Data:   nil,
			Stderr: stderr,
		}, err
	}

	bo = backup_output.BackupOutput{
		Data:   f,
		Stderr: stderr,
	}
	return bo, nil
}
