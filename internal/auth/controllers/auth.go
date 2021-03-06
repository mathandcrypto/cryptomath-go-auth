package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/auth/serializers"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/auth/services"
)

type AuthController struct {
	pbAuth.UnimplementedAuthServiceServer
	authService	*services.AuthService
	refreshSessionSerializer *serializers.RefreshSessionSerializer
}

func (s *AuthController) CreateAccessSession(ctx context.Context,
	req *pbAuth.CreateAccessSessionRequest) (*pbAuth.CreateAccessSessionResponse, error) {
	_, _ = s.authService.DeleteAccessSession(ctx, req.UserId)

	accessSecret, err := s.authService.CreateAccessSession(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create access session: %v", err))
	}

	refreshSecret, err := s.authService.CreateRefreshSession(ctx, req.UserId, accessSecret, req.Ip, req.UserAgent)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create refresh session: %v", err))
	}

	return &pbAuth.CreateAccessSessionResponse{
		AccessSecret:  string(accessSecret),
		RefreshSecret: string(refreshSecret),
	}, nil
}

func (s *AuthController) ValidateAccessSession(ctx context.Context,
	req *pbAuth.ValidateAccessSessionRequest) (*pbAuth.ValidateAccessSessionResponse, error) {
	isSessionExists, err := s.authService.ValidateAccessSession(ctx, req.UserId, req.AccessSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to validate access session: %v", err))
	}

	return &pbAuth.ValidateAccessSessionResponse{
		IsSessionExists: isSessionExists,
	}, nil
}

func (s *AuthController) ValidateRefreshSession(ctx context.Context,
	req *pbAuth.ValidateRefreshSessionRequest) (*pbAuth.ValidateRefreshSessionResponse, error) {
	refreshSession, err := s.authService.FindRefreshSession(ctx, req.UserId, req.RefreshSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to validate refresh session: %v", err))
	}

	if refreshSession == nil {
		return &pbAuth.ValidateRefreshSessionResponse{
			IsSessionExpired: false,
			RefreshSession:   nil,
		}, nil
	}

	isSessionExpired := s.authService.CheckRefreshSessionExpiration(refreshSession)

	return &pbAuth.ValidateRefreshSessionResponse{
		IsSessionExpired: isSessionExpired,
		RefreshSession:   s.refreshSessionSerializer.Serialize(refreshSession),
	}, nil
}

func (s *AuthController) DeleteAccessSession(ctx context.Context,
	req *pbAuth.DeleteAccessSessionRequest) (*pbAuth.DeleteAccessSessionResponse, error) {
	accessSecret, err := s.authService.DeleteAccessSession(ctx, req.UserId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete access session: %v", err))
	}

	isSessionDeleted := false
	if len(accessSecret) > 0 {
		isSessionDeleted = true
	}

	return &pbAuth.DeleteAccessSessionResponse{
		IsSessionDeleted: isSessionDeleted,
	}, nil
}

func (s *AuthController) DeleteRefreshSession(ctx context.Context,
	req *pbAuth.DeleteRefreshSessionRequest) (*pbAuth.DeleteRefreshSessionResponse, error) {
	refreshSession, err := s.authService.DeleteRefreshSession(ctx, req.UserId, req.RefreshSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete refresh session: %v", err))
	}

	return &pbAuth.DeleteRefreshSessionResponse{
		RefreshSession: s.refreshSessionSerializer.Serialize(refreshSession),
	}, nil
}

func (s *AuthController) DeleteAllUserSessions(ctx context.Context,
	req *pbAuth.DeleteAllUserSessionsRequest) (*pbAuth.DeleteAllUserSessionsResponse, error) {
	deletedRefreshSessionsCount, err := s.authService.DeleteAllUserSessions(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete all user sessions: %v", err))
	}

	return &pbAuth.DeleteAllUserSessionsResponse{
		DeletedRefreshSessionsCount: deletedRefreshSessionsCount,
	}, nil
}

func NewAuthController(authConf *authConfig.Config, db *sql.DB, rdb *redis.Client) *AuthController {
	return &AuthController{
		authService: services.NewAuthService(authConf, db, rdb),
		refreshSessionSerializer: &serializers.RefreshSessionSerializer{},
	}
}