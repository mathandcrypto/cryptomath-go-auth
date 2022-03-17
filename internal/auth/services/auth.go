package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/models"
)

type AuthService struct {
	authConf	*authConfig.Config
	db	*sql.DB
	rdb	 *redis.Client
	encryptionService *EncryptionService
}

func (s *AuthService) CreateAccessSession(ctx context.Context, userId int32) ([]byte, error) {
	accessSecret, err := s.encryptionService.GenerateSecret(userId, []byte(""))
	if err != nil {
		return nil, fmt.Errorf("failed to generate user (%d) access secret: %w", userId, err)
	}

	if err = s.rdb.Set(ctx, string(userId), accessSecret, time.Duration(s.authConf.AccessSessionMaxAge)*time.Hour).Err(); err != nil {
		return nil, fmt.Errorf("failed to create user (%d) new access session: %w", userId, err)
	}

	return accessSecret, nil
}

func (s *AuthService) CreateRefreshSession(ctx context.Context, userId int32, accessSecret []byte, ip, userAgent string) ([]byte, error) {
	refreshSecret, err := s.encryptionService.GenerateSecret(userId, accessSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user (%d) refresh secret: %w", userId, err)
	}

	refreshSessionsCount, err := models.RefreshSessions(models.RefreshSessionWhere.UserID.EQ(int(userId))).Count(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count user (%d) refresh sessions: %w", userId, err)
	}

	if refreshSessionsCount > s.authConf.MaxRefreshSessions {
		_, err = models.RefreshSessions(models.RefreshSessionWhere.UserID.EQ(int(userId))).DeleteAll(ctx, s.db)
		if err != nil {
			return nil, fmt.Errorf("failed to delete user (%d) old refresh sessions: %w", userId, err)
		}
	}

	refreshSession := models.RefreshSession{
		RefreshSecret: string(refreshSecret),
		UserID: int(userId),
		IP: ip,
		UserAgent: userAgent,
	}

	if err = refreshSession.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, fmt.Errorf("failed to create user (%d) refresh session: %w", userId, err)
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

func (s *AuthService) FindRefreshSession(ctx context.Context, userId int32, refreshSecret string) (*models.RefreshSession, error) {
	refreshSession, err := models.RefreshSessions(
		models.RefreshSessionWhere.RefreshSecret.EQ(refreshSecret),
		models.RefreshSessionWhere.UserID.EQ(int(userId)),
	).One(ctx, s.db)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to find user (%d) refresh session: %w", userId, err)
	}

	return refreshSession, nil
}

func (s *AuthService) CheckRefreshSessionExpiration(refreshSession *models.RefreshSession) bool {
	return refreshSession.CreatedAt.Before(s.authConf.RefreshSessionExpirationDate())
}

func (s *AuthService) DeleteAccessSession(ctx context.Context, userId int32) ([]byte, error) {
	val, err := s.rdb.GetDel(ctx, string(userId)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to delete user (%d) access session: %w", userId, err)
	}

	return []byte(val), nil
}

func (s *AuthService) DeleteRefreshSession(ctx context.Context, userId int32, refreshSecret string) (*models.RefreshSession, error) {
	refreshSession, err := s.FindRefreshSession(ctx, userId, refreshSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to find user (%d) refresh session to delete a session: %w", userId, err)
	}

	_, err = refreshSession.Delete(ctx, s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user (%d) refresh session: %w", userId, err)
	}

	return refreshSession, nil
}

func (s *AuthService) DeleteAllUserSessions(ctx context.Context, userId int32) (int64, error) {
	_, err := s.DeleteAccessSession(ctx, userId)
	if err != nil {
		return 0, err
	}

	deletedSessionsCount, err := models.RefreshSessions(models.RefreshSessionWhere.UserID.EQ(int(userId))).DeleteAll(ctx, s.db)
	if err != nil {
		return 0, fmt.Errorf("failed to delete all user (%d) refresh sessions: %w", userId, err)
	}

	return deletedSessionsCount, nil
}

func NewAuthService(authConf *authConfig.Config, db *sql.DB, rdb *redis.Client) *AuthService {
	return &AuthService{
		authConf: authConf,
		db: db,
		rdb: rdb,
		encryptionService: &EncryptionService{},
	}
}