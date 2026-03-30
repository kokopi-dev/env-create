package services

import (
	"fmt"
	"os"
	"os/exec"
)

type ScpService struct{}

func NewScpService() *ScpService {
	return &ScpService{}
}

func (s *ScpService) EnvExists() bool {
	_, err := os.Stat(".env")
	return err == nil
}

func (s *ScpService) Run(username, host, remotePath string) (string, error) {
	scpPath, err := exec.LookPath("scp")
	if err != nil {
		return "", fmt.Errorf("scp not found in PATH: %w", err)
	}

	dest := username + "@" + host + ":" + remotePath
	cmd := exec.Command(scpPath, ".env", dest)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
