package backup_process

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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

func (bp *BackupProcess) GetDataStderr() (io.ReadSeeker, string, error) {
	data, err := bp.GetData()
	if err != nil {
		return nil, "", err
	}
	stderr, err := bp.GetStderr()
	if err != nil {
		return nil, "", err
	}
	return data, stderr, nil
}

func (bp *BackupProcess) BaseExecWait(dir string, bin string, args ...string) error {
	var err error

	cmd := exec.Command(bin, args...)
	cmd.Stdout = bp.Data
	cmd.Stderr = bp.StderrBuff
	cmd.Dir = dir

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
	return bp.BaseExecWait("", bin, args...)
}

func (bp *BackupProcess) ExecDirWait(dir string, bin string, args ...string) error {
	return bp.BaseExecWait(dir, bin, args...)
}
