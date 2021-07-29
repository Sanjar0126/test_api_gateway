package main

import (
	"github.com/Sanjar0126/test_api_gateway/api"
	"github.com/Sanjar0126/test_api_gateway/config"
	"github.com/Sanjar0126/test_api_gateway/pkg/grpc_client"
	"github.com/Sanjar0126/test_api_gateway/pkg/logger"
)

var (
	log        logger.Logger
	cfg        config.Config
	grpcClient *grpc_client.GrpcClient
	err        error
)

func initDeps() {
	var err error
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "api_gateway")

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	server.Run(cfg.HTTPPort)
}
