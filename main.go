package main

import (
	"auth/repository"
	"auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	if err := repository.Open(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{ // need for port
		AllowCredentials: true, // for cookie
	}))

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Internal Server Error: %v", err)
	}
}
