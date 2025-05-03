package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"github.com/irfhakeem/go-fiber-clean-starter/service"
)

type (
	AuthController interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
		Verify(ctx *fiber.Ctx) error
	}

	authController struct {
		as service.AuthService
	}
)

func NewAuthController(as service.AuthService) AuthController {
	return &authController{
		as: as,
	}
}

func (c *authController) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_BODY, err.Error()),
		)
	}

	user, err := c.as.Register(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_REGISTER_USER, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		utils.SuccessResponse(dto.SUCCESS_REGISTER_USER, user),
	)
}

func (c *authController) Login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_BODY, err.Error()),
		)
	}

	token, err := c.as.Login(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_LOGIN_USER, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.SuccessResponse(dto.SUCCESS_LOGIN_USER, token),
	)
}

func (c *authController) Verify(ctx *fiber.Ctx) error {
	var req dto.VerifyRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_BODY, err.Error()),
		)
	}

	if err := c.as.Verify(ctx.Context(), req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_VERIFY_USER, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.SuccessResponse(dto.SUCCESS_VERIFY_USER, nil),
	)
}
