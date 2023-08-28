package main

import (
	"PilotSoul/dynamic_service/src/infrastructure"

	"github.com/gofiber/fiber/v2"

	_ "PilotSoul/dynamic_service/src/docs"

	"github.com/gofiber/swagger"
)

// @title Dynamic service
// @version 1.0
// @description Документация для  сервиса, хранящего пользователя и сегменты, в которых он состоит.
// @host 127.0.0.1:3000
func main() {
	infrastructure.ConnectDB()
	app := fiber.New()
	setupRoutes(app)
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Listen(":3000")
}
