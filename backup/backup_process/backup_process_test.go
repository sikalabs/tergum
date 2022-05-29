package backup_process

import (
	"os/exec"
	"testing"
)

func TestStderrBuffer(t *testing.T) {
	bp := BackupProcess{}
	bp.Init()

	const MESSAGE = "helloworld"

	cmd := exec.Command(
		"echo", "-n", MESSAGE,
	)
	cmd.Stdout = bp.StderrBuff
	cmd.Start()
	cmd.Wait()

	stderrString, err := bp.GetStderr()
	if err != nil {
		t.Error(err)
	}
	if stderrString != MESSAGE {
		t.Errorf(`output is "%s" but has to be "%s"`, stderrString, MESSAGE)
	}
}
