package backup_process_utils

import (
	"io"

	"github.com/sikalabs/tergum/backup_process"
)

func BackupProcessExecToFile(bin string, args ...string) (io.ReadSeeker, string, error) {
	var err error
	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return nil, "", err
	}
	err = bp.ExecWait(
		bin,
		args...,
	)
	if err != nil {
		stderr, _ := bp.GetStderr()
		return nil, stderr, err
	}
	return bp.GetDataStderr()
}

func BackupProcessExecEnvToFile(env []string, bin string, args ...string) (io.ReadSeeker, string, error) {
	var err error
	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return nil, "", err
	}
	err = bp.ExecEnvWait(
		env,
		bin,
		args...,
	)
	if err != nil {
		stderr, _ := bp.GetStderr()
		return nil, stderr, err
	}
	return bp.GetDataStderr()
}