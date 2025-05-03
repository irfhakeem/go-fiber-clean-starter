package main

import (
	"log"
	"os"

	"github.com/irfhakeem/go-fiber-clean-starter/config"
	"github.com/irfhakeem/go-fiber-clean-starter/controller"
	"github.com/irfhakeem/go-fiber-clean-starter/entity"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/command"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"github.com/irfhakeem/go-fiber-clean-starter/middleware"
	"github.com/irfhakeem/go-fiber-clean-starter/rabbitmq"
	"github.com/irfhakeem/go-fiber-clean-starter/repository"
	"github.com/irfhakeem/go-fiber-clean-starter/routes"
	"github.com/irfhakeem/go-fiber-clean-starter/service"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

const (
	RabbitMQURL = "amqp://guest:guest@localhost:5672/"
)

func args(db *gorm.DB) bool {
	if len(os.Args) > 1 {
		flag := command.DatabaseCommand(db)
		if !flag {
			return false
		}
	}

	return true
}

func main() {
	db := config.ConnectDatabase()
	defer config.CloseDatabase(db)

	if !args(db) {
		return
	}

	conn, ch := rabbitmq.InitRabbitMQ()
	defer config.CloseConnectionRabbitMQ(conn, ch)

	go rabbitmq.ConsumeAll()

	app := fiber.New()
	app.Use(middleware.Cors())

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Status:  true,
			Message: "Server is healthy",
		})
	})

	// Dependency Injection (Service, Repository, Controller)
	jwtService := service.NewJwtService()

	userRepo := repository.NewBaseRepository[entity.User](db)

	authService := service.NewAuthService(jwtService, userRepo)
	userService := service.NewUserService(userRepo)

	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)

	routes.Auth(app, authController)
	routes.User(app, jwtService, userController)

	app.Static("/uploads", "./uploads")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := app.Listen(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
