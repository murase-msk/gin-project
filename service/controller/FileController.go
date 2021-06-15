//ファイル処理関係
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//アップロードされたファイルの解析
func FileAnalysisAction(c *gin.Context) {
	var e error
	//アップロードファイル取得
	data1, e := c.FormFile("image-input")
	if e != nil {
		//fmt.Println(e)
		defer c.JSON(http.StatusOK, gin.H{"message": e})
		return
	}
	//ファイル保存
	e = c.SaveUploadedFile(data1, "savedFile/"+data1.Filename)
	if e != nil {
		//fmt.Println(e)
		defer c.JSON(http.StatusOK, gin.H{"message": e})
		return
	}

	defer c.JSON(http.StatusOK, gin.H{
		"message": "index message",
	})

}
