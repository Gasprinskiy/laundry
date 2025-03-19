package config

import (
	"fmt"
	"os"
)

// Config структура для хранения переменных окружения
type Config struct {
	DbUrl string
	Port  string
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	return &Config{
		DbUrl: os.Getenv("DATABASE_URL"),
		Port:  fmt.Sprintf(":%s", os.Getenv("PORT")),
	}
}
