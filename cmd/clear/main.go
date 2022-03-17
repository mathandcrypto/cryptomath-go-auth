package main

import (
	"context"

	authConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/auth"
	dbConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/db"
	authJobs "github.com/mathandcrypto/cryptomath-go-auth/internal/auth/jobs"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/common/logger"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/providers"
)

func main()  {
	//	Init context
	ctx := context.Background()

	// Init logger
	l := logger.CreateLogger("auth-clear").WithContext(ctx)

	//	Init database
	dbConf, err := dbConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load database config")
	}

	db, err := providers.NewDBProvider(dbConf)
	if err != nil {
		l.WithError(err).Fatal("failed init database provider")
	}

	//	Init auth config and clear service
	authConf, err := authConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load auth config")
	}

	//	Start clear expired sessions
	clearJob := authJobs.NewClearJob(authConf, db, l)

	clearJob.ClearExpiredSessions(ctx)

	if err = db.Close(); err != nil {
		l.WithError(err).Fatal("failed to close database connection")
	}
}
