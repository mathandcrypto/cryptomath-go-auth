package serializers

import (
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mathandcrypto/cryptomath-go-auth/internal/common/interfaces"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/models"
)

type RefreshSessionSerializer struct {
	interfaces.BaseSerializer
}

func (s *RefreshSessionSerializer) Serialize(refreshSession *models.RefreshSession) *pbAuth.RefreshSession {
	return &pbAuth.RefreshSession{
		Ip:        refreshSession.IP,
		UserAgent: refreshSession.UserAgent,
		CreatedAt: timestamppb.New(refreshSession.CreatedAt),
	}
}
