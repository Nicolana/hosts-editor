package server

import (
	"github.com/Nicolana/hosts-editor/src/server/router"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
