package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBSource	 string
	JWTSecretKey string
	APIPort		 string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return  nil, fmt.Errorf("failed to load .env: %w", err)
	}

	cfg := Config {
		DBSource: os.Getenv("DB_SOURCE"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		APIPort: os.Getenv("API_PORT"),
	}

	return &cfg, nil
}