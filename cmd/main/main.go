package main

import (
	"github.com/fentezi/translete-word/database"
	"github.com/fentezi/translete-word/internal/config"
	"github.com/fentezi/translete-word/internal/server"
	"github.com/fentezi/translete-word/internal/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.MustSetupLogger(cfg.Env)

	db := database.NewRedisDatabase(cfg)

	server.NewGinServer(cfg, db, logger).Start()

}
