package config

import "time"

type Config struct {
	Port         string
	ShutDownTime time.Duration // Tempo para as conexoes terminarem antes de encerrar o servidor
	ReadTimeout  time.Duration // Tempo limite para leitura de uma requisicao
}

func (cfg *Config) Load() {
	*cfg = Config{
		Port:         "8080",
		ShutDownTime: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
}
