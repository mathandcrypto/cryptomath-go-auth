package jobs

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/auth/models"
)

type ClearExpiredSessions struct {
	cron.Job
	ctx     context.Context
	authCfg *authConfig.Config
	db      *gorm.DB
	l       *logrus.Logger
}

func (j *ClearExpiredSessions) Run() {
	tx := j.db.WithContext(j.ctx).Where("created_at < ?", j.authCfg.RefreshSessionExpirationDate()).Delete(&models.RefreshSession{})
	if tx.Error != nil {
		j.l.Error(fmt.Sprintf("failed to clear expired refresh sessions: %v", tx.Error))
		return
	}

	j.l.Infof("deleted %d expired refresh sessions", tx.RowsAffected)
}

func NewClearExpiredSessionsJob(ctx context.Context, cr *cron.Cron,
	authCfg *authConfig.Config, db *gorm.DB, l *logrus.Logger) (cron.EntryID, error) {
	job := ClearExpiredSessions{
		ctx:     ctx,
		authCfg: authCfg,
		db:      db,
		l:       l,
	}

	jobId, err := cr.AddJob("0 * * * *", &job)
	if err != nil {
		return 0, fmt.Errorf("failed to add clear expired sessions job: %w", err)
	}

	return jobId, nil
}
