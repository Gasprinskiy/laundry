package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config структура для хранения переменных окружения
type Config struct {
	DbUrl        string
	Port         string
	RedisPort    string
	RedisPass    string
	RedisTtl     int
	ClientUrl    string
	ClinetDevUrl string
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	redisTtl, err := strconv.Atoi(os.Getenv("REDIS_TTL"))
	if err != nil {
		log.Panic("Could not conver jwt secret lifetime: ", err)
	}

	return &Config{
		DbUrl:        os.Getenv("DATABASE_URL"),
		RedisPass:    os.Getenv("REDIS_PASSWORD"),
		ClientUrl:    os.Getenv("CLIENT_URL"),
		ClinetDevUrl: os.Getenv("CLEINT_DEV_URL"),
		Port:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		RedisPort:    fmt.Sprintf("redis:%s", os.Getenv("REDIS_PORT")),
		RedisTtl:     redisTtl,
	}
}
