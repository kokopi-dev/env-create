package services

import (
	"os"
	"path/filepath"
)

type ProjectNameService struct {
	DefaultProjectName string
}

func NewProjectNameService() *ProjectNameService {
	return &ProjectNameService{DefaultProjectName: "my-project"}
}

func (s *ProjectNameService) GetProjectName() string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = s.DefaultProjectName
	}
	defaultName := filepath.Base(cwd)
	return defaultName
}
