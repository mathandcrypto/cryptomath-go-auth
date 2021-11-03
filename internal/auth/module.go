package auth

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	authControllers "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/controllers"
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Init(grpcServer *grpc.Server, rdb *redis.Client, db *gorm.DB) error {
	authController, err := authControllers.NewAuthController(rdb, db)
	if err != nil {
		return fmt.Errorf("failed to init auth controller: %w", err)
	}

	pbAuth.RegisterAuthServiceServer(grpcServer, authController)

	return nil
}
