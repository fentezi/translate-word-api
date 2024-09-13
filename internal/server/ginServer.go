package server

import (
	"fmt"
	"log/slog"

	"github.com/fentezi/translete-word/database"
	"github.com/fentezi/translete-word/internal/config"
	"github.com/fentezi/translete-word/internal/handlers"
	"github.com/fentezi/translete-word/internal/repositories"
	"github.com/fentezi/translete-word/internal/services"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app    *gin.Engine
	db     database.Database
	conf   *config.Config
	logger *slog.Logger
}

func NewGinServer(conf *config.Config, db database.Database, logger *slog.Logger) Server {
	app := gin.Default()

	return &ginServer{
		app:    app,
		db:     db,
		conf:   conf,
		logger: logger,
	}
}

func (g *ginServer) Start() {

	g.app.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	g.initializeWordHttpHandler()

	serverURL := fmt.Sprintf("%s:%s", g.conf.Server.Host, g.conf.Server.Port)
	g.app.Run(serverURL)
}

func (g *ginServer) initializeWordHttpHandler() {
	workRedisRepository := repositories.NewWordRedisRepository(g.db)

	wordService := services.NewCacheService(workRedisRepository, g.logger)

	wordHttpHandler := handlers.NewWordHttpHandler(wordService)

	wordRouters := g.app.Group("/api")
	{
		wordRouters.GET("/word", wordHttpHandler.WordTranslate)
	}
}
