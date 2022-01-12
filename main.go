package main

import (
	"auth/repository"
	"auth/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	if err := repository.Open(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	app := fiber.New()
	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Internal Server Error: %v", err)
	}
}
