package server

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.Run()
}
