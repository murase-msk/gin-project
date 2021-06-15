package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexIndexAction(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "index message",
	})
}
