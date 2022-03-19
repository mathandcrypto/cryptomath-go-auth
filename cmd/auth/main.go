package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"

	appConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/app"
	dbConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/db"
	redisConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/redis"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/common/logger"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/providers"
)

func main() {
	//	Init context
	ctx, cancel := context.WithCancel(context.Background())

	// Init logger
	l := logger.CreateLogger("auth").WithContext(ctx)

	//	Init database
	dbConf, err := dbConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load database config")
	}

	db, err := providers.NewDBProvider(dbConf)
	if err != nil {
		l.WithError(err).Fatal("failed init database provider")
	}

	//	Init redis
	redisConf, err := redisConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load redis config")
	}

	rdb, err := providers.NewRedisProvider(ctx, redisConf)
	if err != nil {
		l.WithError(err).Fatal("failed to init redis provider")
	}

	//	Init app
	appConf, err := appConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load application config")
	}

	lis, err := net.Listen("tcp", appConf.Address())
	if err != nil {
		l.WithError(err).Fatal("failed to listen local network address")
	}

	var grpcOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(grpcOptions...)

	if err = auth.Init(grpcServer, db, rdb); err != nil {
		l.WithError(err).Fatal("failed to init auth module")
	}

	//	Subscribe to system signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		stop := <-signalChan

		l.WithField("signal", stop).Info("waiting for all processes to stop")
		cancel()

		if err = db.Close(); err != nil {
			l.WithError(err).Fatal("failed to close database connection")
		}

		grpcServer.GracefulStop()
	}()

	l.Info(fmt.Sprintf("starting grpc server on: %s", appConf.Address()))

	if err = grpcServer.Serve(lis); err != nil {
		l.WithError(err).Fatal("failed to serve grpc server")
	}

	wg.Wait()
	l.Info("service stopped")
}
