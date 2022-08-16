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
		// ini配置
		frp.GET("/list", controller.GetForwards)
		frp.POST("/update", controller.UpdateForward)
		frp.POST("/delete", controller.DelForward)
		frp.POST("/add", controller.AddForward)

		// frp客户端
		frp.GET("/start", controller.StartFrp)
		frp.GET("/stop", controller.StopFrp)
		frp.GET("/restart", controller.RestartFrp)
		frp.GET("/status", controller.GetFrpStatus)

		// 日志文件
		frp.GET("/log", controller.GetLog)

		// frp 服务器端设置
		frp.GET("/server", controller.GetServerConfig)
		frp.POST("/server", controller.UpdateServerConfig)
	}
}
