package kubernetes_tls_secret

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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

func (s KubernetesTLSSecret) Backup() (io.ReadSeeker, error) {
	var err error

	outputFile, err := os.CreateTemp("", "tergum-dump-k8s-tls-secret-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(outputFile.Name())

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

	cmd := exec.Command(
		"kubectl",
		args...,
	)
	cmd.Stdout = outputFile

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	outputFile.Seek(0, 0)
	return outputFile, nil
}
