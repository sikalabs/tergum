package kubernetes_tls_secret

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/backup_process_utils"
)

type KubernetesTLSSecret struct {
	Server     string `yaml:"Server"`
	Token      string `yaml:"Token"`
	Namespace  string `yaml:"Namespace"`
	SecretName string `yaml:"SecretName"`
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

func (s KubernetesTLSSecret) Backup() (io.ReadSeeker, string, error) {
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
