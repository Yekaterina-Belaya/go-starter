package routes

import (
	"go-starter/app-1/backend/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.HomePage)
}