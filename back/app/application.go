package app

import (
	"ideal-journey/controller"
	"log"
)

func Init() {
	handlers := controller.NewHandler()
	app := handlers.Init()
	log.Fatal(app.Listen(":8080"))
}
