package handlers

import "github.com/gin-gonic/gin"

type WordHandler interface {
	WordTranslate(c *gin.Context)
}
