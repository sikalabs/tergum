package backup_process_utils

import (
	"io"

	"github.com/sikalabs/tergum/backup/backup_process"
)

func BackupProcessExecToFile(bin string, args ...string) (io.ReadSeeker, string, error) {
	var err error
	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return nil, "", err
	}
	bp.ExecWait(
		bin,
		args...,
	)
	if err != nil {
		return nil, "", err
	}
	return bp.GetDataStderr()
}
