/*
コントローラーパッケージ

関数命名規則　NameController.go -> func NameMethodAction()
*/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murase-msk/gin-project/service"
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

	var labelList []string
	imageProc := new(service.ImageProcess)
	//labelList = imageProc.GetLabelFromImage("savedFile/" + data1.Filename)
	//jsonVal, _ := json.Marshal(labelList)

	imageProc.GetTextFromImage("savedFile/" + data1.Filename)

	defer c.JSON(http.StatusOK, gin.H{
		"message": labelList,
	})

}
