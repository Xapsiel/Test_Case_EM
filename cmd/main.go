package main

import (
	"fmt"
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/Xapsiel/EffectiveMobile/internal/handler"
	"github.com/Xapsiel/EffectiveMobile/internal/repository"
	"github.com/Xapsiel/EffectiveMobile/internal/service"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Тут должен быть лог")
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		fmt.Println("Тут должен быть лог", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(EffectiveMobile.Server)
	if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {

		fmt.Println("Тут должен быть лог")
	}
}
