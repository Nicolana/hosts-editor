package server

import (
	"net/http"

	"github.com/Nicolana/hosts-editor/src/server/router"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	router.InitRouter(r)
	// global.SystemInit()
	r.LoadHTMLFiles("./web/dist/index.html")
	r.Static("/assets", "./web/dist/assets")
	r.StaticFile("vite.svg", "./web/dist")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":8022")
}
