package services

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type ScpConfig struct {
	Username   string `json:"username"`
	Host       string `json:"host"`
	RemotePath string `json:"remote_path"`
}

type ConfigIndex struct {
	Default  *ScpConfig            `json:"default,omitempty"`
	Projects map[string]*ScpConfig `json:"projects"`
}

type ConfigService struct {
	cacheDir  string
	indexPath string
}

func NewConfigService() (*ConfigService, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cacheDir := filepath.Join(home, ".cache", "env-create")
	return &ConfigService{
		cacheDir:  cacheDir,
		indexPath: filepath.Join(cacheDir, "index.json"),
	}, nil
}

func (s *ConfigService) CacheExists() bool {
	_, err := os.Stat(s.cacheDir)
	return err == nil
}

func (s *ConfigService) CreateCacheDir() error {
	return os.MkdirAll(s.cacheDir, 0755)
}

func (s *ConfigService) LoadConfig(projectPath string) (*ScpConfig, error) {
	data, err := os.ReadFile(s.indexPath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var idx ConfigIndex
	if err := json.Unmarshal(data, &idx); err != nil {
		return nil, err
	}

	if cfg, ok := idx.Projects[projectPath]; ok {
		return cfg, nil
	}
	return idx.Default, nil
}

func (s *ConfigService) SaveConfig(projectPath string, cfg ScpConfig) error {
	var idx ConfigIndex
	data, err := os.ReadFile(s.indexPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err == nil {
		if err := json.Unmarshal(data, &idx); err != nil {
			return err
		}
	}

	if idx.Projects == nil {
		idx.Projects = make(map[string]*ScpConfig)
	}
	cfgCopy := cfg
	idx.Projects[projectPath] = &cfgCopy
	if idx.Default == nil {
		idx.Default = &cfgCopy
	}

	out, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.indexPath, out, 0644)
}
