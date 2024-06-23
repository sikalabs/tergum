package kubernetes

import (
	"fmt"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process_utils"
)

type Kubernetes struct {
	Server    string `yaml:"Server" json:"Server,omitempty"`
	Token     string `yaml:"Token" json:"Token,omitempty"`
	Namespace string `yaml:"Namespace" json:"Namespace,omitempty"`
	Resource  string `yaml:"Resource" json:"Resource,omitempty"`
	Name      string `yaml:"Name" json:"Name,omitempty"`
}

func (s Kubernetes) Validate() error {
	if s.Server == "" {
		return fmt.Errorf("Kubernetes need to have a Server")
	}
	if s.Token == "" {
		return fmt.Errorf("Kubernetes need to have a Token")
	}
	if s.Resource == "" {
		return fmt.Errorf("Kubernetes need to have a Resource")
	}
	return nil
}

func (s Kubernetes) Backup() (backup_output.BackupOutput, error) {
	args := []string{
		"--server", s.Server,
		"--insecure-skip-tls-verify=true",
		"--token", s.Token,
		"get",
		s.Resource,
		"-o", "yaml",
	}
	if s.Name != "" {
		args = append(args, s.Name)
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
