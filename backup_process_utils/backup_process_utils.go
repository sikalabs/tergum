package backup_process_utils

import (
	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process"
)

func BackupProcessExecToFile(bin string, args ...string) (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return bo, err
	}
	err = bp.ExecWait(
		bin,
		args...,
	)
	if err != nil {
		stderr, _ := bp.GetStderr()
		bo.Stderr = stderr
		return bo, err
	}
	return bp.GetDataStderr()
}

func BackupProcessExecEnvToFile(env []string, bin string, args ...string) (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return bo, err
	}
	err = bp.ExecEnvWait(
		env,
		bin,
		args...,
	)
	if err != nil {
		stderr, _ := bp.GetStderr()
		bo.Stderr = stderr
		return bo, err
	}
	return bp.GetDataStderr()
}
