package router

import (
	"github.com/Nicolana/hosts-editor/src/server/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		api.GET("/list", controller.GetList)
		api.POST("/update", controller.UpdateRow)
		api.POST("/delete", controller.DeleteRow)
		api.POST("/add", controller.AddRow)
	}
}
