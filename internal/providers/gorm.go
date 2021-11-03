package providers

import (
	"context"
	"fmt"
	databaseConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewGORMProvider(ctx context.Context, config *databaseConfig.Config) (*gorm.DB, error) {
	//	TODO: add custom gorm-logrus module
	gormLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel: logger.Silent,
		IgnoreRecordNotFoundError:	true,
		Colorful:	false,
	})

	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db.WithContext(ctx), nil
}
