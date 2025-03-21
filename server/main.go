package main

import (
	"laundry/config"
	external "laundry/external/ginapi"
	"laundry/internal/repository/rimport"
	"laundry/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Получение переменных окружения
	conf := config.NewConfig()

	// Подключение к базе
	db, err := sqlx.Open("postgres", conf.DbUrl)
	if err != nil {
		log.Panic("Ошибка при подключении к базе данных: ", err)
	}
	// Закрыть соеденение к базе перед выходом из функции main
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Panic("Ошибка при пинге базы данных: ", err)
	}

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     conf.RedisPort,
	// 	Password: conf.RedisPass,
	// 	DB:       0,
	// })

	r := gin.Default()

	repo := rimport.NewRepositoryImports()
	servicesUsecase := usecase.NewServicesUsecase(repo, db)
	ordersUsecase := usecase.NewOrdersUsecase(repo, db)

	external.RegiserServicesExternal(servicesUsecase, r)
	external.RegiserOrdersExternal(ordersUsecase, r)

	r.Run(conf.Port)
}
