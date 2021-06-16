package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func IndexIndexAction(c *gin.Context) {

	//環境変数読み込み
	err := godotenv.Load("./.env_dev")
	if err != nil {
		// .env読めなかった場合の処理
		c.JSON(200, gin.H{"message": err})
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":  "index message",
		"message2": os.Getenv("ENV"),
	})
}
