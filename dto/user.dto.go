package dto

import (
	"errors"
	"mime/multipart"
	"time"

	"github.com/irfhakeem/go-fiber-clean-starter/helpers/constants"
)

const (
	FAILED_GET_ALL_USERS  = "Failed to get all users"
	FAILED_GET_USER_BY_ID = "Failed to get user by id"
	FAILED_CREATE_USER    = "Failed to create user"
	FAILED_UPDATE_USER    = "Failed to update user"
	FAILED_DELETE_USER    = "Failed to delete user"

	SUCCESS_GET_ALL_USERS  = "Success to get all users"
	SUCCESS_GET_USER_BY_ID = "Success to get user by id"
	SUCCESS_CREATE_USER    = "Success to create user"
	SUCCESS_UPDATE_USER    = "Success to update user"
	SUCCESS_DELETE_USER    = "Success to delete user"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type (
	UserResponse struct {
		ID        int64              `json:"id"`
		Email     string             `json:"email"`
		Name      string             `json:"name"`
		Gender    constants.Gender   `json:"gender"`
		Role      constants.UserRole `json:"role"`
		Avatar    string             `json:"avatar"`
		CreatedAt time.Time          `json:"created_at"`
		UpdatedAt time.Time          `json:"updated_at"`
	}

	UserPaginationResponse struct {
		Users []UserResponse `json:"data"`
		PaginationResponse
	}

	UserCreateRequest struct {
		Email string `json:"email" form:"email" validate:"required,email"`
		Name  string `json:"name" form:"name" validate:"required"`
	}

	UserUpdateRequest struct {
		Email    string                `json:"email" form:"email" validate:"omitempty,email"`
		Name     string                `json:"name" form:"name" validate:"omitempty"`
		Password string                `json:"password" form:"password" validate:"omitempty"`
		Gender   constants.Gender      `json:"gender" form:"password" validate:"omitempty"`
		Avatar   *multipart.FileHeader `json:"avatar" form:"avatar" validate:"omitempty"`
	}
)
