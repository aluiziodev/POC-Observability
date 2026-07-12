package config

import (
	"os"
	"time"
)

type Config struct {
	Port         string
	ShutDownTime time.Duration // Tempo para as conexoes terminarem antes de encerrar o servidor
	ReadTimeout  time.Duration // Tempo limite para leitura de uma requisicao
}

func Load() Config {
	return Config{
		Port:         getEnvOrDefault("API_PORT", "8080"),
		ShutDownTime: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
