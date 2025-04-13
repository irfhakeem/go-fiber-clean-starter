package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irfhakeem/go-fiber-clean-starter/controller"
	"github.com/irfhakeem/go-fiber-clean-starter/middleware"
	"github.com/irfhakeem/go-fiber-clean-starter/service"
)

func User(
	app *fiber.App,
	jwtService service.JWTService,
	c controller.UserController,
) {
	routes := app.Group("/api/user")

	routes.Get("/me", middleware.Authorize(jwtService), c.Me)
	routes.Get("/", middleware.Authorize(jwtService), c.FindAllUsers)
	routes.Get("/:id", middleware.Authorize(jwtService), c.FindUserByID)

	routes.Post("/", middleware.Authorize(jwtService), c.CreateUser)

	routes.Put("/:id", middleware.Authorize(jwtService), c.UpdateUser)

	routes.Delete("/:id", middleware.Authorize(jwtService), c.DeleteUser)
}
