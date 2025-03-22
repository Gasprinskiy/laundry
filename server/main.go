package main

import (
	"context"
	"laundry/config"
	external "laundry/external/ginapi"
	"laundry/internal/repository/rimport"
	"laundry/internal/usecase"
	"laundry/redisclient"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Получение переменных окружения
	conf := config.NewConfig()

	// Подключение к базе
	db, err := sqlx.Open("postgres", conf.DbUrl)
	if err != nil {
		log.Panic("Ошибка при подключении к базе данных: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Panic("Ошибка при пинге базы данных: ", err)
	}

	// Закрыть соеденение к базе перед выходом из функции main
	defer db.Close()

	// Подлключение к redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.RedisPort,
		Password: conf.RedisPass,
		DB:       0,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Panic("Ошибка при подключении пинге redis: ", err)
	}

	// Закрыть соеденение к базе перед выходом из функции main
	defer rdb.Close()

	r := gin.Default()

	redisCleint := redisclient.NewRedisClient(rdb, conf)
	repo := rimport.NewRepositoryImports()
	servicesUsecase := usecase.NewServicesUsecase(repo, db)
	ordersUsecase := usecase.NewOrdersUsecase(repo, db)

	external.RegiserServicesExternal(servicesUsecase, r)
	external.RegiserOrdersExternal(ordersUsecase, redisCleint, r)

	r.Run(conf.Port)
}
