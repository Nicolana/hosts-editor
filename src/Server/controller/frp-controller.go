package controller

import (
	"github.com/Nicolana/hosts-editor/src/frp"
	"github.com/Nicolana/hosts-editor/src/server/config"
	"github.com/Nicolana/hosts-editor/src/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddForward 添加新的映射
func AddForward(c *gin.Context) {
	var updateData models.FrpSectionType
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "字段不正确", "data": gin.H{}, "code": config.ErrorCode})
		return
	}
	//item, err := f.Editor.UpdateLineByIndex(updateData.Row-1, updateData.Hosts, updateData.Ip, updateData.Comments)
	item, err := frp.FrpcIni.UpdateSecByName(updateData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": gin.H{}, "code": config.ErrorCode})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "修改成功", "data": item.ToMap(), "code": config.SuccessCode})
}

// DelForward 删除映射
func DelForward(c *gin.Context) {}

// UpdateForward 修改端口映射
func UpdateForward(c *gin.Context) {
}

// GetForwards 获取端口映射列表
func GetForwards(c *gin.Context) {
	err := frp.FrpcIni.LoadFile()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "无法加载数据", "data": gin.H{}, "code": config.ErrorCode})
	}
	search := c.DefaultQuery("search", "") // shortcut for c.Request.URL.Query().Get("lastname")
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "20")
	pageNum, _ := strconv.Atoi(page)
	sizeNum, _ := strconv.Atoi(size)

	c.JSON(http.StatusOK, gin.H{"message": "加载成功", "data": frp.FrpcIni.GetList(pageNum, sizeNum, search), "code": config.SuccessCode})
}
