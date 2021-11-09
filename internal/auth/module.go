package auth

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	authControllers "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/controllers"
	authJobs "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/jobs"
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Init(ctx context.Context, cr *cron.Cron, grpcServer *grpc.Server, rdb *redis.Client, db *gorm.DB, l *logrus.Logger) error {
	authCfg, err := authConfig.New()
	if err != nil {
		return fmt.Errorf("failed to load auth config: %w", err)
	}

	clearExpiredSessionsJobId, err := authJobs.NewClearExpiredSessionsJob(ctx, cr, authCfg, db, l)
	if err != nil {
		return fmt.Errorf("failed to init clear expired sessions job: %w", err)
	}
	l.Info(fmt.Sprintf("added clear expired sessions job (id: %d)", clearExpiredSessionsJobId))

	authController := authControllers.NewAuthController(rdb, db, authCfg)

	pbAuth.RegisterAuthServiceServer(grpcServer, authController)

	return nil
}
