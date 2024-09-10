package main

import (
	"github.com/fentezi/translete-word/internal/config"
	"github.com/fentezi/translete-word/internal/handlers"
	red "github.com/fentezi/translete-word/internal/repositories/redis"
	"github.com/fentezi/translete-word/internal/routers"
	"github.com/fentezi/translete-word/internal/services"
	"github.com/fentezi/translete-word/internal/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	_ = logger.MustSetupLogger(cfg.Env)

	redis := config.NewRedis(cfg)

	repo := red.NewRedisRepository(redis)

	service := services.NewService(repo)

	handlers := handlers.NewHandlers(service)

	routers := routers.NewRouters(handlers)

	routers.InitRouter().Run(cfg.Server.Host + ":" + cfg.Server.Port)

}
