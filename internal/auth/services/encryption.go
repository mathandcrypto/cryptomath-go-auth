package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type EncryptionService struct{}

func (s *EncryptionService) GenerateSecret(userId int32, extra []byte) ([]byte, error) {
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid (v4): %w", err)
	}

	now := time.Now()
	plain := []byte(fmt.Sprintf("%d%s%s%d", userId, extra, uuidV4, now.Unix()))

	hash, err := bcrypt.GenerateFromPassword(plain, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash plain secret: %w", err)
	}

	return hash, nil
}