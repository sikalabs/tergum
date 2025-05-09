package proxmox_local_vm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process"
)

type ProxmoxLocalVMSoure struct {
	VMID int `yaml:"VMID" json:"VMID,omitempty"`
}

func (s ProxmoxLocalVMSoure) Validate() error {
	if s.VMID == 0 {
		return fmt.Errorf("ProxmoxLocalVMSoure need to have a VMID")
	}
	return nil
}

func (s ProxmoxLocalVMSoure) Backup() (backup_output.BackupOutput, error) {
	var bo backup_output.BackupOutput
	var err error

	bp := backup_process.BackupProcess{}
	bp.Init()

	tempDir, err := os.MkdirTemp("", "tergum-vzdump-"+strconv.Itoa(s.VMID)+"-*")
	if err != nil {
		return backup_output.BackupOutput{
			Data: nil,
		}, err
	}
	defer os.RemoveAll(tempDir) // Clean up when done

	backupFilePath := tempDir + "/tergum-vzdump-qemu-" + strconv.Itoa(s.VMID) + ".zst"

	args := []string{
		strconv.Itoa(s.VMID),
		"--dumpdir", tempDir,
		"--mode", "snapshot",
		"--compress", "zstd",
	}

	fmt.Println("vzdump", args)
	err = bp.ExecWait("vzdump", args...)
	stderr, _ := bp.GetStderr()
	if err != nil {
		return backup_output.BackupOutput{
			Data:   nil,
			Stderr: stderr,
		}, err
	}

	err = bp.ExecWait("sh", "-c", "mv "+tempDir+"/vzdump-qemu-"+strconv.Itoa(s.VMID)+"-*.zst "+backupFilePath)
	stderr, _ = bp.GetStderr()
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
