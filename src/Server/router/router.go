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
	hosts := api.Group("/hosts")
	{
		hosts.GET("/list", controller.GetList)
		hosts.POST("/update", controller.UpdateRow)
		hosts.POST("/delete", controller.DeleteRow)
		hosts.POST("/add", controller.AddRow)
	}
	frp := api.Group("/frp")
	{
		frp.GET("/list", controller.GetForwards)
		frp.POST("/update", controller.UpdateForward)
		frp.POST("/delete", controller.DelForward)
		frp.POST("/add", controller.AddForward)
	}
}
