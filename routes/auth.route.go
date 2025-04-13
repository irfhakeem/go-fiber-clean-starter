package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irfhakeem/go-fiber-clean-starter/controller"
)

func Auth(
	app *fiber.App,
	c controller.AuthController,
) {
	routes := app.Group("/api/auth")

	routes.Post("/register", c.Register)
	routes.Post("/login", c.Login)
}
