package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server ServerConfig
	RZD    RZDConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type RZDConfig struct {
	BaseURL string
	Timeout time.Duration
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		},
		RZD: RZDConfig{
			BaseURL: getEnv("RZD_API_URL", "http://localhost:8000"),
			Timeout: time.Duration(getEnvAsInt("RZD_TIMEOUT", 10)) * time.Second,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
