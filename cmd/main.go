package main

import (
	"github.com/Xapsiel/EffectiveMobile/internal/handler"
	"github.com/Xapsiel/EffectiveMobile/internal/models"
	"github.com/Xapsiel/EffectiveMobile/internal/repository"
	"github.com/Xapsiel/EffectiveMobile/internal/service"
	"github.com/Xapsiel/EffectiveMobile/pkg/log"
	"github.com/joho/godotenv"
	"os"
)

//	@title			Songs API
// 	@version 1.0
//	@description	This is an API for managing songs.

//	@host			localhost:8080
//	@BasePath		/

func main() {
	logService := log.NewLogService(os.Stdout, "debug")
	if err := godotenv.Load(); err != nil {
		log.Logger.Info(log.MakeLog("Ошибка загрузки .env", err))
	}
	log.Logger.Debug(log.MakeLog("Загрузки переменных окружения", nil))
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Logger.Info(log.MakeLog("Ошибка создания объекта repository", err))
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	services.Log = logService
	handlers := handler.NewHandler(services)
	srv := new(models.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Logger.Info(log.MakeLog("Ошибка запуска сервера", err))
	}
}
