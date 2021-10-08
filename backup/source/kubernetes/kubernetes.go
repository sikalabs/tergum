package kubernetes

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Kubernetes struct {
	Server    string `yaml:"Server"`
	Token     string `yaml:"Token"`
	Namespace string `yaml:"Namespace"`
	Resource  string `yaml:"Resource"`
	Name      string `yaml:"Name"`
}

func (s Kubernetes) Validate() error {
	if s.Server == "" {
		return fmt.Errorf("Kubernetes need to have a Server")
	}
	if s.Namespace == "" {
		return fmt.Errorf("Kubernetes need to have a Namespace")
	}
	if s.Namespace == "" {
		return fmt.Errorf("Kubernetes need to have a Namespace")
	}
	if s.Resource == "" {
		return fmt.Errorf("Kubernetes need to have a Resource")
	}
	return nil
}

func (s Kubernetes) Backup() (io.ReadSeeker, error) {
	var err error

	outputFile, err := os.CreateTemp("", "tergum-dump-k8s-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(outputFile.Name())

	args := []string{
		"--server", s.Server,
		"--insecure-skip-tls-verify=true",
		"--token", s.Token,
		"--namespace", s.Namespace,
		"get",
		s.Resource,
		"-o", "yaml",
	}
	if s.Name != "" {
		args = append(args, s.Name)
	}

	cmd := exec.Command(
		"kubectl",
		args...,
	)
	cmd.Stdout = outputFile

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	cmd.Wait()

	outputFile.Seek(0, 0)
	return outputFile, nil
}
