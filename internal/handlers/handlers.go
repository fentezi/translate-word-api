package handlers

import (
	"net/http"

	"github.com/fentezi/translete-word/internal/services"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Services *services.Service
}

func NewHandlers(service *services.Service) *Handlers {
	return &Handlers{
		Services: service,
	}
}

func (h *Handlers) GetTranslateWord(c *gin.Context) {
	word := c.Query("word")

	translate, err := h.Services.GetTranslateWord(word)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"translate": translate})

}
