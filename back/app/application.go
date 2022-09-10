package app

import (
	"ideal-journey/clients/cassandra"
	"ideal-journey/clients/logger"
	"ideal-journey/controller"
	"ideal-journey/usecase"
	"ideal-journey/usecase/repo"
	"log"
)

func Init() {
	cassandra := cassandra.NewClient()
	repos := repo.NewRepositories(cassandra)
	services := usecase.NewServices(usecase.Dependecies{
		Repos: repos,
	})
	handlers := controller.NewHandler(services)
	app := handlers.Init()

	logger.Info("Server started")
	log.Fatal(app.Listen("localho.st:8080"))
}
