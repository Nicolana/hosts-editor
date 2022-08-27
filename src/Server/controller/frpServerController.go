package controller

import (
	"fmt"
	"net/http"

	"github.com/Nicolana/hosts-editor/src/server/config"
	"github.com/Nicolana/hosts-editor/src/server/global"
	"github.com/Nicolana/hosts-editor/src/server/models"
	"github.com/gin-gonic/gin"
)

type FrpServerController struct{}

var db = global.Config.Db

// 新增服务器
func (server *FrpServerController) NewServer(c *gin.Context) {
	var payload models.FrpServer
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "字段不正确", "data": gin.H{}, "code": config.ErrorCode})
		return
	}
	item := db.Omit("ID", "CreatedAt", "UpdatedAt", "DeletedAt").Create(&payload)
	c.JSON(http.StatusOK, gin.H{"message": "数据插入失败", "data": item, "code": config.SuccessCode})
}

// 删除服务器
func (server *FrpServerController) DeleteServer(c *gin.Context) {
	fmt.Println(c)
}

// 获取服务器列表
func (server *FrpServerController) GetServerList(c *gin.Context) {
	fmt.Println(c)
}

// 更新服务器配置
func (server *FrpServerController) UpdateServer(c *gin.Context) {
	fmt.Println(c)
}
