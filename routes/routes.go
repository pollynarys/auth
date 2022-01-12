package routes

import (
	"auth/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)
}
