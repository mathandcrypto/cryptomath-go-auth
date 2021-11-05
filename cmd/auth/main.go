package main

import (
	"context"
	"fmt"
	appConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/app"
	databaseConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/database"
	redisConfig "github.com/mathandcrypto/cryptomath-go-auth/configs/redis"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/auth"
	"github.com/mathandcrypto/cryptomath-go-auth/internal/providers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func setupLogger(ctx context.Context) *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	l.WithContext(ctx)

	return l
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l := setupLogger(ctx)

	//	Init database
	dbConfig, err := databaseConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load database config")
	}

	db, err := providers.NewGORMProvider(ctx, l, dbConfig)
	if err != nil {
		l.WithError(err).Fatal("failed init database provider")
	}

	//	Init redis
	redisCfg, err := redisConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load redis config")
	}

	rdb, err := providers.NewRedisProvider(ctx, redisCfg)
	if err != nil {
		l.WithError(err).Fatal("failed to init redis provider")
	}

	//	Init app
	appCfg, err := appConfig.New()
	if err != nil {
		l.WithError(err).Fatal("failed to load app config")
	}

	lis, err := net.Listen("tcp", appCfg.Address())
	if err != nil {
		l.WithError(err).Fatal("failed to listen")
	}

	var grpcOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(grpcOptions...)

	if err = auth.Init(grpcServer, rdb, db); err != nil {
		l.WithError(err).Fatal("failed to init auth module")
	}

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
		grpcServer.GracefulStop()
		rdb.FlushDB(ctx)
	}()

	l.Info(fmt.Sprintf("starting grpc server on: %s", appCfg.Address()))
	if err = grpcServer.Serve(lis); err != nil {
		l.WithError(err).Fatal("failed to serve grpc server")
	}

	wg.Wait()
	l.Info("service stopped")
}