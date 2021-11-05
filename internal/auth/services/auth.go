package authServices

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	authModels "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/models"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	rdb *redis.Client
	db *gorm.DB
	authCfg *authConfig.Config
	encryptionService *EncryptionService
}

func (s *AuthService) CreateAccessSession(ctx context.Context, userId int32) ([]byte, error) {
	accessSecret, err := s.encryptionService.GenerateSecret(userId, []byte(""))
	if err != nil {
		return nil, fmt.Errorf("failed to generate user (%d) access secret: %w", userId, err)
	}

	if err = s.rdb.Set(ctx, string(userId), accessSecret, time.Duration(s.authCfg.AccessSessionMaxAge) * time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("failed to create user (%d) new access session: %w", userId, err)
	}

	return accessSecret, nil
}

func (s *AuthService) CreateRefreshSession(ctx context.Context, userId int32, accessSecret []byte, ip string, userAgent string) ([]byte, error) {
	refreshSecret, err := s.encryptionService.GenerateSecret(userId, accessSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user (%d) refresh secret: %w", userId, err)
	}

	var refreshSessionsCount int64
	tx := s.db.WithContext(ctx).Model(&authModels.RefreshSession{}).Where(&authModels.RefreshSession{UserId: userId}).Count(&refreshSessionsCount)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to count user (%d) refresh sessions: %w", userId, tx.Error)
	}

	if refreshSessionsCount > s.authCfg.MaxRefreshSessions {
		tx = s.db.WithContext(ctx).Where(&authModels.RefreshSession{UserId: userId}).Delete(&authModels.RefreshSession{})
		if tx.Error != nil {
			return nil, fmt.Errorf("failed to delete user (%d) old refresh sessions: %w", userId, tx.Error)
		}
	}

	refreshSession := &authModels.RefreshSession{
		RefreshSecret: string(refreshSecret),
		UserId: userId,
		IP: ip,
		UserAgent: userAgent,
	}

	tx = s.db.WithContext(ctx).Create(&refreshSession)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to create user (%d) new refresh session: %w", userId, tx.Error)
	}

	return refreshSecret, nil
}

func (s *AuthService) ValidateAccessSession(ctx context.Context, userId int32, accessSecret string) (bool, error) {
	redisAccessSecret, err := s.rdb.Get(ctx, string(userId)).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to get user (%d) access session: %w", userId, err)
	}

	return accessSecret == redisAccessSecret, nil
}

func (s *AuthService) ValidateRefreshSession(ctx context.Context, userId int32, refreshSecret string) (bool, bool, error) {
	type APIRefreshSession struct {
		CreatedAt time.Time
	}

	var result APIRefreshSession
	tx := s.db.WithContext(ctx).Model(&authModels.RefreshSession{}).First(&result).Where(&authModels.RefreshSession{
		UserId: userId,
		RefreshSecret: refreshSecret,
	})
	if tx.Error != nil {
		return false, false, fmt.Errorf("failed to get user (%d) refresh session: %w", userId, tx.Error)
	}

	return false, false, nil
}

func (s *AuthService) DeleteAccessSession(ctx context.Context, userId int32) ([]byte, error) {
	val, err := s.rdb.GetDel(ctx, string(userId)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to delete access session: %w", err)
	}

	return []byte(val), nil
}

func NewAuthService(rdb *redis.Client, db *gorm.DB, authCfg *authConfig.Config) *AuthService {
	return &AuthService{
		rdb: rdb,
		db: db,
		authCfg: authCfg,
		encryptionService: NewEncryptionService(),
	}
}
