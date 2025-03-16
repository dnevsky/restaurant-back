package rest

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoot(api *gin.Engine) {
	api.GET("/", h.root)
}

func (h *Handler) root(c *gin.Context) {
	c.JSON(200, gin.H{"response": "я живой"})
}
