package main

import (
	"gitlab.udevs.io/delever/delever_customer_api_gateway/api"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/config"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/grpc_client"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
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
