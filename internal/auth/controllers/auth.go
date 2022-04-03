package controllers

import (
	"context"
	"database/sql"

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

func (c *AuthController) CreateAccessSession(ctx context.Context,
	req *pbAuth.CreateAccessSessionRequest) (*pbAuth.CreateAccessSessionResponse, error) {
	_, _ = c.authService.DeleteAccessSession(ctx, req.UserId)

	accessSecret, err := c.authService.CreateAccessSession(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access session: %v", err)
	}

	refreshSecret, err := c.authService.CreateRefreshSession(ctx, req.UserId, accessSecret, req.Ip, req.UserAgent)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh session: %v", err)
	}

	return &pbAuth.CreateAccessSessionResponse{
		AccessSecret:  string(accessSecret),
		RefreshSecret: string(refreshSecret),
	}, nil
}

func (c *AuthController) ValidateAccessSession(ctx context.Context,
	req *pbAuth.ValidateAccessSessionRequest) (*pbAuth.ValidateAccessSessionResponse, error) {
	isSessionExists, err := c.authService.ValidateAccessSession(ctx, req.UserId, req.AccessSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to validate access session: %v", err)
	}

	return &pbAuth.ValidateAccessSessionResponse{
		IsSessionExists: isSessionExists,
	}, nil
}

func (c *AuthController) ValidateRefreshSession(ctx context.Context,
	req *pbAuth.ValidateRefreshSessionRequest) (*pbAuth.ValidateRefreshSessionResponse, error) {
	refreshSession, err := c.authService.FindRefreshSession(ctx, req.UserId, req.RefreshSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to validate refresh session: %v", err)
	}

	if refreshSession == nil {
		return &pbAuth.ValidateRefreshSessionResponse{
			IsSessionExpired: false,
			RefreshSession:   nil,
		}, nil
	}

	isSessionExpired := c.authService.CheckRefreshSessionExpiration(refreshSession)

	return &pbAuth.ValidateRefreshSessionResponse{
		IsSessionExpired: isSessionExpired,
		RefreshSession:   c.refreshSessionSerializer.Serialize(refreshSession),
	}, nil
}

func (c *AuthController) DeleteAccessSession(ctx context.Context,
	req *pbAuth.DeleteAccessSessionRequest) (*pbAuth.DeleteAccessSessionResponse, error) {
	accessSecret, err := c.authService.DeleteAccessSession(ctx, req.UserId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete access session: %v", err)
	}

	isSessionDeleted := false
	if len(accessSecret) > 0 {
		isSessionDeleted = true
	}

	return &pbAuth.DeleteAccessSessionResponse{
		IsSessionDeleted: isSessionDeleted,
	}, nil
}

func (c *AuthController) DeleteRefreshSession(ctx context.Context,
	req *pbAuth.DeleteRefreshSessionRequest) (*pbAuth.DeleteRefreshSessionResponse, error) {
	refreshSession, err := c.authService.DeleteRefreshSession(ctx, req.UserId, req.RefreshSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete refresh session: %v", err)
	}

	return &pbAuth.DeleteRefreshSessionResponse{
		RefreshSession: c.refreshSessionSerializer.Serialize(refreshSession),
	}, nil
}

func (c *AuthController) DeleteAllUserSessions(ctx context.Context,
	req *pbAuth.DeleteAllUserSessionsRequest) (*pbAuth.DeleteAllUserSessionsResponse, error) {
	deletedRefreshSessionsCount, err := c.authService.DeleteAllUserSessions(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete all user sessions: %v", err)
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