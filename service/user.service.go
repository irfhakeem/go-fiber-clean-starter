package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/entity"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"github.com/irfhakeem/go-fiber-clean-starter/repository"
)

type (
	UserService interface {
		FindAllUsers(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		FindUserByID(ctx context.Context, id int64) (dto.UserResponse, error)
		CreateUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		UpdateUser(ctx context.Context, req dto.UserUpdateRequest, id int64) (dto.UserResponse, error)
		DeleteUser(ctx context.Context, id int64) error
	}

	userService struct {
		us repository.IBaseRepository[entity.User]
	}
)

func NewUserService(us repository.IBaseRepository[entity.User]) UserService {
	return &userService{
		us: us,
	}
}

func formatUserResponse(u *entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Gender:    u.Gender,
		Role:      u.Role,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (s *userService) FindAllUsers(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error) {
	users, err := s.us.FindAll(ctx, nil, req, "name LIKE ? OR email LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	if err != nil {
		return dto.UserPaginationResponse{}, err
	}

	var userResponses []dto.UserResponse
	for _, user := range users.Data {
		userResponses = append(userResponses, formatUserResponse(&user))
	}

	return dto.UserPaginationResponse{
		Users:              userResponses,
		PaginationResponse: users.PaginationResponse,
	}, nil
}

func (s *userService) FindUserByID(ctx context.Context, id int64) (dto.UserResponse, error) {
	user, err := s.us.FindByID(ctx, nil, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return formatUserResponse(user), nil
}

func (s *userService) CreateUser(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	user, err := s.us.Create(ctx, nil, &entity.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Email + req.Name[0:1],
	})
	if err != nil {
		return dto.UserResponse{}, err
	}

	return formatUserResponse(user), nil
}

func (s *userService) UpdateUser(ctx context.Context, req dto.UserUpdateRequest, id int64) (dto.UserResponse, error) {
	user, err := s.us.FindByID(ctx, nil, id)
	if err != nil {
		return dto.UserResponse{}, dto.ErrUserNotFound
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return dto.UserResponse{}, err
		}
		user.Password = hashedPassword
	}

	if req.Avatar != nil {
		fileName := fmt.Sprintf(
			"%s/%s-%s",
			"avatar",
			strconv.FormatInt(user.ID, 10),
			strings.ReplaceAll(req.Avatar.Filename, " ", "-"),
		)

		if err := utils.Uploads(*req.Avatar, fileName); err != nil {
			return dto.UserResponse{}, err
		}
		user.Avatar = fileName
	}

	updatedUser, err := s.us.Update(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrUpdateUser
	}

	return formatUserResponse(updatedUser), nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	err := s.us.Delete(ctx, nil, id)
	if err != nil {
		return err
	}

	return nil
}
