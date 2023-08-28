package controllers

import (
	"PilotSoul/dynamic_service/src/domain"
	"PilotSoul/dynamic_service/src/infrastructure"

	"github.com/gofiber/fiber/v2"
)

// CreateUser func creates user.
// @Description Создание пользователя.
// @Accept json
// @Param input body domain.User true "Пользователь"
// @Success 200
// @Router /create_user [post]
func CreateUser(c *fiber.Ctx) error {
	// Создание пользователя
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if user.Name == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Name cannot be empty",
		})
	}
	infrastructure.DB.Db.Create(&user)
	return c.Status(200).JSON(user)
}
