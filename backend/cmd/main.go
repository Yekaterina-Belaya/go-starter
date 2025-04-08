package main

import (
	"go-starter/app-1/backend/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main(){
app := fiber.New()


app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("hello")
} )

routes.SetupRoutes(app)

app.Listen(":3000")


}

