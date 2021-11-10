package authServices

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	authModels "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/models"
)

type AuthService struct {
	rdb               *redis.Client
	db                *gorm.DB
	authCfg           *authConfig.Config
	encryptionService *EncryptionService
}

func (s *AuthService) CreateAccessSession(ctx context.Context, userId int32) ([]byte, error) {
	accessSecret, err := s.encryptionService.GenerateSecret(userId, []byte(""))
	if err != nil {
		return nil, fmt.Errorf("failed to generate user (%d) access secret: %w", userId, err)
	}

	if err = s.rdb.Set(ctx, string(userId), accessSecret, time.Duration(s.authCfg.AccessSessionMaxAge)*time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("failed to create user (%d) new access session: %w", userId, err)
	}

	return accessSecret, nil
}

func (s *AuthService) CreateRefreshSession(ctx context.Context,
	userId int32, accessSecret []byte, ip string, userAgent string) ([]byte, error) {
	refreshSecret, err := s.encryptionService.GenerateSecret(userId, accessSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user (%d) refresh secret: %w", userId, err)
	}

	var refreshSessionsCount int64
	tx := s.db.WithContext(ctx).
		Model(&authModels.RefreshSession{}).Where(&authModels.RefreshSession{UserId: userId}).Count(&refreshSessionsCount)
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
		UserId:        userId,
		IP:            ip,
		UserAgent:     userAgent,
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

func (s *AuthService) FindRefreshSession(ctx context.Context, userId int32, refreshSecret string) (*authModels.RefreshSession, error) {
	var result authModels.RefreshSession
	tx := s.db.WithContext(ctx).Take(&result).Where(&authModels.RefreshSession{
		UserId:        userId,
		RefreshSecret: refreshSecret,
	})
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to find user (%d) refresh session: %w", userId, tx.Error)
	}

	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return &result, nil
}

func (s *AuthService) CheckRefreshSessionExpiration(refreshSession *authModels.RefreshSession) bool {
	return refreshSession.CreatedAt.Before(s.authCfg.RefreshSessionExpirationDate())
}

func (s *AuthService) DeleteAccessSession(ctx context.Context, userId int32) ([]byte, error) {
	val, err := s.rdb.GetDel(ctx, string(userId)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to delete user (%d) access session: %w", userId, err)
	}

	return []byte(val), nil
}

func (s *AuthService) DeleteRefreshSession(ctx context.Context, userId int32, refreshSecret string) (*authModels.RefreshSession, error) {
	refreshSession, err := s.FindRefreshSession(ctx, userId, refreshSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to find user (%d) refresh session to deletion: %w", userId, err)
	}

	tx := s.db.WithContext(ctx).Delete(refreshSession)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to delete user (%d) refresh session: %w", userId, err)
	}

	return refreshSession, nil
}

func (s *AuthService) DeleteAllUserSessions(ctx context.Context, userId int32) error {
	_, err := s.DeleteAccessSession(ctx, userId)
	if err != nil {
		return err
	}

	tx := s.db.WithContext(ctx).Delete(&authModels.RefreshSession{
		UserId: userId,
	})
	if tx.Error != nil {
		return fmt.Errorf("failed to delete all user (%d) refresh sessions: %w", userId, tx.Error)
	}

	return nil
}

func NewAuthService(rdb *redis.Client, db *gorm.DB, authCfg *authConfig.Config) *AuthService {
	return &AuthService{
		rdb:               rdb,
		db:                db,
		authCfg:           authCfg,
		encryptionService: NewEncryptionService(),
	}
}
