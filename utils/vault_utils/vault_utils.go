package vault_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const KubernetesSATokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"

func GetVaultTokenFromKubernetesSA(addr, role, authPath, tokenFile string) (string, error) {
	if authPath == "" {
		authPath = "kubernetes"
	}
	if tokenFile == "" {
		tokenFile = KubernetesSATokenPath
	}

	saToken, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", fmt.Errorf("failed to read Kubernetes SA token at %s: %w", tokenFile, err)
	}

	body, err := json.Marshal(map[string]string{
		"role": role,
		"jwt":  string(saToken),
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal vault login request: %w", err)
	}

	url := fmt.Sprintf("%s/v1/auth/%s/login", addr, authPath)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("vault kubernetes login request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read vault login response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("vault kubernetes login failed (status %d): %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		Auth struct {
			ClientToken string `json:"client_token"`
		} `json:"auth"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse vault login response: %w", err)
	}
	if result.Auth.ClientToken == "" {
		return "", fmt.Errorf("vault login response did not contain a client_token")
	}

	return result.Auth.ClientToken, nil
}
