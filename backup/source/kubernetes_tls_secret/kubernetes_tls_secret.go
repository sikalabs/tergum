package kubernetes_tls_secret

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type KubernetesTLSSecret struct {
	Server     string `yaml:"Server" json:"Server,omitempty"`
	Token      string `yaml:"Token" json:"Token,omitempty"`
	Namespace  string `yaml:"Namespace" json:"Namespace,omitempty"`
	SecretName string `yaml:"SecretName" json:"SecretName,omitempty"`
}

func (s KubernetesTLSSecret) Validate() error {
	if s.Server == "" {
		return fmt.Errorf("KubernetesTLSSecret need to have a Server")
	}
	if s.Token == "" {
		return fmt.Errorf("KubernetesTLSSecret need to have a Token")
	}
	return nil
}

func (s KubernetesTLSSecret) Backup() (backup_output.BackupOutput, error) {
	args := []string{
		"--server", s.Server,
		"--insecure-skip-tls-verify=true",
		"--token", s.Token,
		"get",
		"secrets",
		"--field-selector", "type=kubernetes.io/tls",
		"-o", "yaml",
	}
	if s.SecretName != "" {
		args = append(args, s.SecretName)
	}
	if s.Namespace != "" {
		args = append(args, "--namespace", s.Namespace)
	} else {
		args = append(args, "--all-namespaces")
	}

	return backup_process_utils.BackupProcessExecToFile(
		"kubectl",
		args...,
	)
}
