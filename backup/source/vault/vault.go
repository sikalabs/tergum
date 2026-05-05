package vault

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
	"github.com/sikalabs/tergum/utils/vault_utils"
)

type VaultSource struct {
	Addr               string            `yaml:"Addr" json:"Addr,omitempty"`
	Token              string            `yaml:"Token" json:"Token,omitempty"`
	KubernetesRole     string            `yaml:"KubernetesRole" json:"KubernetesRole,omitempty"`
	KubernetesAuthPath string            `yaml:"KubernetesAuthPath" json:"KubernetesAuthPath,omitempty"`
	Headers            map[string]string `yaml:"Headers" json:"Headers,omitempty"`
}

func (s VaultSource) getToken() (string, error) {
	if s.Token != "" {
		return s.Token, nil
	}
	return vault_utils.GetVaultTokenFromKubernetesSA(s.Addr, s.KubernetesRole, s.KubernetesAuthPath)
}

func (s VaultSource) Validate() error {
	if s.Addr == "" {
		return fmt.Errorf("VaultSource need to have a Addr")
	}
	if s.Token == "" && s.KubernetesRole == "" {
		return fmt.Errorf("VaultSource requires Token or KubernetesRole")
	}
	return nil
}

func (s VaultSource) Backup() (backup_output.BackupOutput, error) {
	token, err := s.getToken()
	if err != nil {
		return backup_output.BackupOutput{}, err
	}
	return backup_process_utils.BackupProcessHttpGetWithToken(
		s.Addr+"/v1/sys/storage/raft/snapshot",
		"X-Vault-Token",
		token,
		s.Headers,
	)
}
