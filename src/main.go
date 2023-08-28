package main

import (
	"PilotSoul/dynamic_service/src/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {
	infrastructure.ConnectDB()
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
