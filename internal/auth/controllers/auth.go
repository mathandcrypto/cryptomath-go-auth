package authControllers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	authServices "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/services"
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthController struct {
	pbAuth.AuthServiceServer
	authService *authServices.AuthService
}

func (s *AuthController) CreateAccessSession(ctx context.Context, req *pbAuth.CreateAccessSessionRequest) (*pbAuth.CreateAccessSessionResponse, error) {
	_, _ = s.authService.DeleteAccessSession(ctx, req.UserId)

	accessSecret, err := s.authService.CreateAccessSession(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create access session: %s", err))
	}

	refreshSecret, err := s.authService.CreateRefreshSession(ctx, req.UserId, accessSecret, req.Ip, req.UserAgent)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create refresh session: %s", err))
	}

	return &pbAuth.CreateAccessSessionResponse{
		AccessSecret: string(accessSecret),
		RefreshSecret: string(refreshSecret),
	}, nil
}

func (s* AuthController) ValidateAccessSession(ctx context.Context, req *pbAuth.ValidateAccessSessionRequest) (*pbAuth.ValidateAccessSessionResponse, error) {
	isSessionExists, err := s.authService.ValidateAccessSession(ctx, req.UserId, req.AccessSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to validate access session: %s", err))
	}

	return &pbAuth.ValidateAccessSessionResponse{
		IsSessionExists: isSessionExists,
	}, nil
}

func NewAuthController(rdb *redis.Client, db *gorm.DB) (*AuthController, error) {
	authCfg, err := authConfig.New()

	if err != nil {
		return nil, fmt.Errorf("failed to load auth config: %w", err)
	}

	return &AuthController{
		authService: authServices.NewAuthService(rdb, db, authCfg),
	}, nil
}
