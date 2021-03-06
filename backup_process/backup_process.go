package backup_process

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/sikalabs/tergum/backup_output"
)

type BackupProcess struct {
	StderrBuff *strings.Builder
	Data       io.ReadWriteSeeker
}

func (bp *BackupProcess) Init() {
	bp.StderrBuff = new(strings.Builder)
}

func (bp *BackupProcess) InitDataTempFile() error {
	f, err := os.CreateTemp("", "tergum-")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	bp.Data = f
	return nil
}

func (bp BackupProcess) GetStderr() (string, error) {
	if bp.StderrBuff == nil {
		return "", fmt.Errorf("StderrBuff is nil")
	}
	return bp.StderrBuff.String(), nil
}

func (bp *BackupProcess) GetData() (io.ReadSeeker, error) {
	_, err := bp.Data.Seek(0, 0)
	return bp.Data, err
}

func (bp *BackupProcess) GetDataStderr() (backup_output.BackupOutput, error) {
	bo := backup_output.BackupOutput{}
	data, err := bp.GetData()
	bo.Data = data
	if err != nil {
		return bo, err
	}
	stderr, err := bp.GetStderr()
	bo.Stderr = stderr
	if err != nil {
		return bo, err
	}
	return bo, nil
}

func (bp *BackupProcess) BaseExecWait(env []string, dir string, bin string, args ...string) error {
	var err error

	cmd := exec.Command(bin, args...)
	cmd.Stdout = bp.Data
	cmd.Stderr = bp.StderrBuff
	cmd.Dir = dir
	cmd.Env = env

	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (bp *BackupProcess) ExecWait(bin string, args ...string) error {
	return bp.BaseExecWait(nil, "", bin, args...)
}

func (bp *BackupProcess) ExecDirWait(dir string, bin string, args ...string) error {
	return bp.BaseExecWait(nil, dir, bin, args...)
}

func (bp *BackupProcess) ExecEnvWait(env []string, bin string, args ...string) error {
	return bp.BaseExecWait(env, "", bin, args...)
}
