package main

import (
	"template-go-auth-mysql/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	routes.SetupAuth(app);

	app.Listen(":5000")

}