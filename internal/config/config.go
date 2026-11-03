package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	Addr      string
	Port      string
	DbURL     string
	JwtSecret string
	LogLevel  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Unable to load .env - %v", err)
	}
	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		Addr:      getEnv("ADDR", ""),
		Port:      getEnv("PORT", "8080"),
		DbURL:     getRequiredEnv("DB_URL"),
		JwtSecret: getRequiredEnv("JWT_SECRET"),
		LogLevel:  getEnv("LOG_LEVEL", "INFO"),
	}

	if cfg.JwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET must not be empty")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if os.Getenv(key) == "" {
		return fallback
	}
	return os.Getenv(key)
}

func getRequiredEnv(key string) string {
	return os.Getenv(key)
}
