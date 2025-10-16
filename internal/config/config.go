package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBSource     string
	JWTSecretKey string
	APIPort      string
	Environment  string
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := Config{
		DBSource:     getEnv("DB_SOURCE", ""),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", ""),
		APIPort:      getEnv("API_PORT", "8080"),
		Environment:  getEnv("ENV", "development"),
	}

	if cfg.DBSource == "" {
		return nil, fmt.Errorf("DB_SOURCE environment variable is required")
	}

	if cfg.JWTSecretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable is required")
	}

	if len(cfg.JWTSecretKey) < 32 {
		return nil, fmt.Errorf("JWT_SECRET_KEY must be at least 32 characters long")
	}

	return &cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}