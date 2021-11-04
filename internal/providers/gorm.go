package providers

import (
	"context"
	"fmt"
	databaseConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/database"
	"github.com/mathandcrypto/cryptomath-gorm-logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func NewGORMProvider(ctx context.Context, l *logrus.Logger, config *databaseConfig.Config) (*gorm.DB, error) {
	newGormLogger := logger.New(l, logger.Config{
		SlowThreshold: time.Second,
		LogLevel: gormLogger.Silent,
		SkipErrRecordNotFound: true,
	})

	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{
		Logger: newGormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db.WithContext(ctx), nil
}
