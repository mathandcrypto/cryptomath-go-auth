package auth

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"google.golang.org/grpc"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/auth/controllers"
)

func Init(grpcServer *grpc.Server, db *sql.DB, rdb *redis.Client) error {
	authConf, err := authConfig.New()
	if err != nil {
		return fmt.Errorf("failed to load auth config: %w", err)
	}

	authController := controllers.NewAuthController(authConf, db, rdb)

	pbAuth.RegisterAuthServiceServer(grpcServer, authController)

	return nil
}
