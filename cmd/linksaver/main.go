package main

import (
	link "LinkSaver"
	"LinkSaver/pkg/handler"
	"LinkSaver/pkg/repository"
	"LinkSaver/pkg/service"
	"log"
)


func main() {
	repos := new(repository.Repository)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(link.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
