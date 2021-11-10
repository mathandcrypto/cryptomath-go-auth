package authSerializers

import (
	pbAuth "github.com/mathandcrypto/cryptomath-go-proto/auth"
	"google.golang.org/protobuf/types/known/timestamppb"

	authModels "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/models"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/common/interfaces"
)

type RefreshSessionSerializer struct {
	interfaces.BaseSerializer
}

func (s *RefreshSessionSerializer) Serialize(refreshSession *authModels.RefreshSession) *pbAuth.RefreshSession {
	return &pbAuth.RefreshSession{
		Ip:        refreshSession.IP,
		UserAgent: refreshSession.UserAgent,
		CreatedAt: timestamppb.New(refreshSession.CreatedAt),
	}
}

func NewRefreshSessionSerializer() *RefreshSessionSerializer {
	return &RefreshSessionSerializer{}
}