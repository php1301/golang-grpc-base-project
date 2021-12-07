package main

import (
	"context"
	"runtime"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/xdorro/golang-grpc-base-project/internal/config"
	"github.com/xdorro/golang-grpc-base-project/internal/server"
	"github.com/xdorro/golang-grpc-base-project/pkg/client"
	"github.com/xdorro/golang-grpc-base-project/pkg/logger"
	"github.com/xdorro/golang-grpc-base-project/pkg/redis"
)

const (
	defaultShutdownTimeout = 10 * time.Second
)

func init() {
	// Load config environment
	config.NewConfig()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create new logger
	log := logger.NewLogger()
	log.Info(viper.GetString("APP_NAME"),
		zap.String("app-version", viper.GetString("APP_VERSION")),
		zap.String("go-version", runtime.Version()),
	)

	// declare new client
	db := client.NewClient(ctx, log)
	// declare new redis
	rdb := redis.NewRedis(log)

	// create new server
	srv, err := server.NewServer(ctx, log, db, rdb)
	if err != nil {
		log.Panic("server.NewServer()", zap.Error(err))
	}

	// wait for termination signal and register database & http server clean-up operations
	wait := gracefulShutdown(ctx, log, defaultShutdownTimeout, map[string]operation{
		"server": func(ctx context.Context) error {
			return srv.Close()
		},
		"client": func(ctx context.Context) error {
			return db.Close()
		},
		"redis": func(ctx context.Context) error {
			return rdb.Close()
		},
	})

	<-wait
}
