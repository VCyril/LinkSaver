package main

import (
	link "LinkSaver"
	"LinkSaver/pkg/handler"
	"LinkSaver/pkg/repository"
	"LinkSaver/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)


func main() {
	ctx := new(gin.Context)
	if err := initConfig(); err != nil {
		log.Fatalf("Error while initializing config: %s", err.Error())
	}

	repos := new(repository.Repository)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(link.Server)
	defer srv.Shutdown(ctx)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigFile("configs/config.yml")
	return viper.ReadInConfig()
}
