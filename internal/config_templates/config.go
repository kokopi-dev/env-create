package config

import (
	"cmp"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	// dev | prod
	Env string
}

// feel free to exit the app if NewAppConfig returns an error
func NewAppConfig() (*AppConfig, error) {
	// env
	if os.Getenv("ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			slog.Warn("no .env file found")
		} else {
			slog.Info(".env loaded")
		}
	}

	var missingVars []string

	required := []string{
		"ENV",
	}

	for _, key := range required {
		if os.Getenv(key) == "" {
			missingVars = append(missingVars, key)
		}
	}

	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s",
			strings.Join(missingVars, ", "))
	}
	return &AppConfig{
		Env:                 cmp.Or(os.Getenv("ENV"), "dev"),
	}, nil
}

func (c *AppConfig) IsDev() bool {
	return c.Env == "dev"
}

func (c *AppConfig) IsProd() bool {
	return c.Env == "prod"
}
