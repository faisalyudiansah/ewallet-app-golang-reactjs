package main

import (
	"ewallet-server-v2/internal/config"
	"ewallet-server-v2/internal/httpserver"
	"ewallet-server-v2/internal/pkg/logger"
)

func main() {
	cfg := config.InitConfig()

	logger.SetLogrusLogger(cfg)

	httpserver.StartGinHttpServer(cfg)
}
