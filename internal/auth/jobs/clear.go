package jobs

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/models"
)

type ClearJob struct {
	authConf	*authConfig.Config
	db	*sql.DB
	l *logrus.Entry
}

func (j *ClearJob) ClearExpiredSessions(ctx context.Context) {
	deletedSessionsCount, err :=
		models.RefreshSessions(models.RefreshSessionWhere.CreatedAt.LTE(j.authConf.RefreshSessionExpirationDate())).DeleteAll(ctx, j.db)
	if err != nil {
		j.l.Error(fmt.Sprintf("failed to clear expired refresh sessions: %v", err))
		return
	}

	j.l.Infof("deleted %d expired refresh sessions", deletedSessionsCount)
}

func NewClearJob(authConf *authConfig.Config, db *sql.DB, l *logrus.Entry) *ClearJob {
	return &ClearJob{
		authConf: authConf,
		db: db,
		l: l,
	}
}
