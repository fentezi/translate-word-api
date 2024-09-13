package handlers

import (
	"net/http"

	"github.com/fentezi/translete-word/internal/models"
	"github.com/fentezi/translete-word/internal/services"
	"github.com/gin-gonic/gin"
)

type wordHttpHandler struct {
	wordService services.WordService
}

func NewWordHttpHandler(wordService services.WordService) WordHandler {
	return &wordHttpHandler{
		wordService: wordService,
	}
}

func (h *wordHttpHandler) WordTranslate(c *gin.Context) {
	var addWord models.AddWord
	word := c.Query("word")

	addWord.Word = word
	translation, err := h.wordService.WordTranslate(&addWord)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": translation,
	})
}
