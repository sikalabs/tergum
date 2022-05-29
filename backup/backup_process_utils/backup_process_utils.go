package backup_process_utils

import (
	"io"

	"github.com/sikalabs/tergum/backup/backup_process"
)

func BackupProcessExecToFile(bin string, args ...string) (io.ReadSeeker, string, error) {
	bp := backup_process.BackupProcess{}
	bp.Init()
	bp.InitDataTempFile()
	bp.ExecWait(
		bin,
		args...,
	)
	return bp.GetDataStderr()
}
