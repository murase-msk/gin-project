package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloTestAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test HelloController",
	})
}

func HelloIndexAction(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.html", gin.H{
		"message": "index message",
	})
}
