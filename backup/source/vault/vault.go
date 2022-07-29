package vault

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type VaultSource struct {
	Addr  string `yaml:"Addr"`
	Token string `yaml:"Token"`
}

func (s VaultSource) Validate() error {
	if s.Addr == "" {
		return fmt.Errorf("VaultSource need to have a Addr")
	}
	if s.Token == "" {
		return fmt.Errorf("VaultSource need to have a Token")
	}
	return nil
}

func (s VaultSource) Backup() (backup_output.BackupOutput, error) {
	return backup_process_utils.BackupProcessHttpGetWithToken(
		s.Addr+"/v1/sys/storage/raft/snapshot",
		"X-Vault-Token",
		s.Token,
	)
}
