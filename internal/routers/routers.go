package routers

import (
	"github.com/fentezi/translete-word/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	handlers *handlers.Handlers
}

func NewRouters(handlers *handlers.Handlers) *Routers {
	return &Routers{
		handlers: handlers,
	}
}

func (r *Routers) InitRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/word", r.handlers.GetTranslateWord)

	return router
}

func (r *Routers) StartServer(host, port string) {

	address := host + ":" + port

	r.InitRouter().Run(address)
}
