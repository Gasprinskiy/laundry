package config

import (
	"fmt"
	"os"
)

// Config структура для хранения переменных окружения
type Config struct {
	DbUrl     string
	Port      string
	RedisPort string
	RedisPass string
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	return &Config{
		DbUrl:     os.Getenv("DATABASE_URL"),
		RedisPass: os.Getenv("REDIS_PASSWORD"),
		Port:      fmt.Sprintf(":%s", os.Getenv("PORT")),
		RedisPort: fmt.Sprintf("localhost:%s", os.Getenv("REDIS_PORT")),
	}
}
