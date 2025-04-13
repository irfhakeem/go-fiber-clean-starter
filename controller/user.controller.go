package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"github.com/irfhakeem/go-fiber-clean-starter/service"
)

type (
	UserController interface {
		Me(ctx *fiber.Ctx) error
		FindAllUsers(ctx *fiber.Ctx) error
		FindUserByID(ctx *fiber.Ctx) error
		CreateUser(ctx *fiber.Ctx) error
		UpdateUser(ctx *fiber.Ctx) error
		DeleteUser(ctx *fiber.Ctx) error
	}

	userController struct {
		us service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		us: us,
	}
}

func (c *userController) Me(ctx *fiber.Ctx) error {
	userIDStr := ctx.Locals("user_id").(string)
	if userIDStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_HEADER, dto.FAILED_HEADER_IS_MISSING),
		)
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_HEADER, err.Error()),
		)
	}

	user, err := c.us.FindUserByID(ctx.Context(), userID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_USER_BY_ID, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.SuccessResponse(dto.SUCCESS_GET_USER_BY_ID, user),
	)
}

func (c *userController) FindAllUsers(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_QUERY, err.Error()),
		)
	}

	users, err := c.us.FindAllUsers(ctx.Context(), req)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_ALL_USERS, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.Response{
			Status:  true,
			Message: dto.SUCCESS_GET_ALL_USERS,
			Data:    users.Users,
			Meta:    users.PaginationResponse,
		},
	)
}

func (c *userController) FindUserByID(ctx *fiber.Ctx) error {
	userIDStr := ctx.Params("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_PARAMS, err.Error()),
		)
	}

	user, err := c.us.FindUserByID(ctx.Context(), userID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_USER_BY_ID, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.SuccessResponse(dto.SUCCESS_GET_USER_BY_ID, user),
	)
}

func (c *userController) CreateUser(ctx *fiber.Ctx) error {
	var req dto.UserCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_BODY, err.Error()),
		)
	}

	user, err := c.us.CreateUser(ctx.Context(), req)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_CREATE_USER, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		utils.SuccessResponse(dto.SUCCESS_CREATE_USER, user),
	)
}

func (c *userController) UpdateUser(ctx *fiber.Ctx) error {
	userIDStr := ctx.Params("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_PARAMS, err.Error()),
		)
	}

	var req dto.UserUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_BODY, err.Error()),
		)
	}

	user, err := c.us.UpdateUser(ctx.Context(), req, userID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_UPDATE_USER, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.SuccessResponse(dto.SUCCESS_UPDATE_USER, user),
	)
}

func (c *userController) DeleteUser(ctx *fiber.Ctx) error {
	userIDStr := ctx.Params("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_GET_DATA_FROM_PARAMS, err.Error()),
		)
	}

	err = c.us.DeleteUser(ctx.Context(), userID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			utils.FailedResponse(dto.FAILED_DELETE_USER, err.Error()),
		)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
