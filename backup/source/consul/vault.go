package consul

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type ConsulSource struct {
	Addr  string `yaml:"Addr"`
	Token string `yaml:"Token"`
}

func (s ConsulSource) Validate() error {
	if s.Addr == "" {
		return fmt.Errorf("ConsulSource need to have a Addr")
	}
	return nil
}

func (s ConsulSource) Backup() (backup_output.BackupOutput, error) {
	return backup_process_utils.BackupProcessHttpGetWithToken(
		s.Addr+"/v1/snapshot",
		"X-Consul-Token",
		s.Token,
	)
}
