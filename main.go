package main

import (
	"auth/repository"
	"auth/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	repository.Open()

	app := fiber.New()
	routes.Setup(app)

	log.Println("start serv")
	app.Listen(":8000")
}
