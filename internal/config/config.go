package config

import (
	"order-service/internal/utils"
	"time"
)

type Config struct {
	Port         string
	ShutDownTime time.Duration // Tempo para as conexoes terminarem antes de encerrar o servidor
	ReadTimeout  time.Duration // Tempo limite para leitura de uma requisicao
}

func Load() Config {
	return Config{
		Port:         utils.GetEnvOrDefault("API_PORT", "8080"),
		ShutDownTime: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
}
