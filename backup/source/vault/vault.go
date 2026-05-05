package vault

import (
	"fmt"
	"os"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

const k8sSATokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"

type VaultSource struct {
	Addr    string            `yaml:"Addr" json:"Addr,omitempty"`
	Token   string            `yaml:"Token" json:"Token,omitempty"`
	Headers map[string]string `yaml:"Headers" json:"Headers,omitempty"`
}

func (s VaultSource) getToken() (string, error) {
	if s.Token != "" {
		return s.Token, nil
	}
	data, err := os.ReadFile(k8sSATokenPath)
	if err != nil {
		return "", fmt.Errorf("VaultSource requires a Token or a Kubernetes SA token at %s", k8sSATokenPath)
	}
	return string(data), nil
}

func (s VaultSource) Validate() error {
	if s.Addr == "" {
		return fmt.Errorf("VaultSource need to have a Addr")
	}
	// Try to get token from config or Kubernetes SA
	_, err := s.getToken()
	if err != nil {
		return err
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
