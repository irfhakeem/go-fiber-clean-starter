package service

import (
	"context"

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
	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return dto.UserResponse{}, err
		}
		req.Password = hashedPassword
	}

	var avatarPath string
	if req.Avatar != nil {
		avatar, err := utils.Uploads(*req.Avatar, "/avatar")
		if err != nil {
			return dto.UserResponse{}, err
		}
		avatarPath = avatar
	}

	user, err := s.us.Update(ctx, nil, &entity.User{
		ID:       id,
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Gender:   req.Gender,
		Avatar:   avatarPath,
	})
	if err != nil {
		return dto.UserResponse{}, err
	}

	return formatUserResponse(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	err := s.us.Delete(ctx, nil, id)
	if err != nil {
		return err
	}

	return nil
}
