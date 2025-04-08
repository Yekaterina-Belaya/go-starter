package handlers

import (
	"go-starter/app-1/backend/internal/app/models"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("index.html", nil)
}

func ListVacancies(c *fiber.Ctx) error {
	vacancies := models.GetAllVacancies()
	return c.JSON(vacancies)
}

func CreateVacancy(c *fiber.Ctx) error {
    var vacancy models.Vacancy
    if err := c.BodyParser(&vacancy); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    models.CreateVacancy(vacancy)
    return c.Status(201).SendString("Vacancy created")
}