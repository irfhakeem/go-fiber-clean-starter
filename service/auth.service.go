package service

import (
	"context"
	"fmt"
	"strings"
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
		Verify(ctx context.Context, req dto.VerifyRequest) error
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

	txt := fmt.Sprintf("%s_%s", user.Email, time.Now().Add(time.Minute*5).Format("2006-01-02 15:04:05"))

	token, err := utils.GetAESEncrypted(txt)
	if err != nil {
		return dto.UserResponse{}, err
	}

	link := fmt.Sprintf("http://localhost:8080/verify?token=%s", token)

	if err := SendVerificationEmail(ctx, user.Email, link); err != nil {
		return dto.UserResponse{}, err
	}

	return formatUserResponse(user), nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.us.FindFirst(ctx, nil, "email = ?", req.Email)
	if user == nil || err != nil {
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

func (s *authService) Verify(ctx context.Context, req dto.VerifyRequest) error {
	txt, err := utils.GetAESDecrypted(req.Token)
	if err != nil {
		return err
	}

	parts := strings.Split(string(txt), "_")
	if len(parts) < 2 {
		return dto.ErrTokenInvalid
	}

	email := parts[0]
	expiredAt, err := time.Parse("2006-01-02 15:04:05", parts[1])
	if err != nil {
		return err
	}

	if time.Now().After(expiredAt) {
		return dto.ErrTokenExpired
	}

	user, err := s.us.FindFirst(ctx, nil, "email = ?", email)
	if err != nil {
		return dto.ErrUserNotFound
	}

	_, err = s.us.Update(ctx, nil, &entity.User{
		ID:         user.ID,
		IsVerified: true,
	})
	if err != nil {
		return dto.ErrUpdateUser
	}

	return nil
}
