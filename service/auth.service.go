package service

import (
	"context"
	"time"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/entity"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"github.com/irfhakeem/go-fiber-clean-starter/repository"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto.RegisterRequest) (dto.UserResponse, error)
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	}

	authService struct {
		jwt JWTService
		us  repository.IBaseRepository[entity.User]
	}
)

func NewAuthService(jwt JWTService, us repository.IBaseRepository[entity.User]) AuthService {
	return &authService{
		jwt: jwt,
		us:  us,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.UserResponse, error) {
	data := entity.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	user, err := s.us.Create(ctx, nil, &data)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return formatUserResponse(user), nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.us.FindFirst(ctx, nil, "email = ?", req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if user == nil {
		return dto.LoginResponse{}, dto.ErrUserNotFound
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return dto.LoginResponse{}, dto.ErrPasswordMismatch
	}

	accToken, err := s.jwt.GenerateToken(dto.JWTPayload{
		UserId: user.ID,
		Role:   user.Role,
	}, time.Now().Add(time.Hour*7))
	if err != nil {
		return dto.LoginResponse{}, dto.ErrGenerateToken
	}

	return dto.LoginResponse{
		AccessToken: accToken,
	}, nil
}
