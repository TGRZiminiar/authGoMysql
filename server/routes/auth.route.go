package routes

import (
	"template-go-auth-mysql/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuth(app *fiber.App){

	app.Post("/api/register", controllers.Register);
	app.Post("/api/login", controllers.Login)
	app.Get("/api/current-user", controllers.CurrentUser)
}
