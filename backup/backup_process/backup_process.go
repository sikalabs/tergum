package backup_process

import (
	"fmt"
	"strings"
)

type BackupProcess struct {
	StderrBuff *strings.Builder
}

func (bp *BackupProcess) Init() {
	bp.StderrBuff = new(strings.Builder)
}

func (bp BackupProcess) GetStderr() (string, error) {
	if bp.StderrBuff == nil {
		return "", fmt.Errorf("StderrBuff is nil")
	}
	return bp.StderrBuff.String(), nil
}
