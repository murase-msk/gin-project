package main

import (
	"github.com/gin-gonic/gin"

	"github.com/murase-msk/gin-project/service/controller"
)

func main() {
	r := gin.Default()
	//事前にテンプレートをロード 相対パス
	r.LoadHTMLGlob("assets/html/*.html")
	// 静的ファイルのパスを指定
	r.Static("/assets", "./assets")

	//ルーティング
	r.GET("/", controller.IndexIndexAction)
	r.GET("/hello", controller.HelloIndexAction)
	r.POST("file-upload", controller.FileAnalysisAction)
	//r.NoRoute(routes.NoRoute) // どのルーティングにも当てはまらなかった場合に処理

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
