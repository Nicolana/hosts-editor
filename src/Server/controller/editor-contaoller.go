package controller

import (
	hostseditor "github.com/Nicolana/hosts-editor/src/hosts-editor"
	"github.com/Nicolana/hosts-editor/src/server/config"
	"github.com/Nicolana/hosts-editor/src/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(c *gin.Context) {
	_, err := hostseditor.Editor.LoadFile()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "无法加载数据", "data": gin.H{}, "code": config.ErrorCode})
	}
	c.JSON(http.StatusOK, hostseditor.Editor.GetLines(1, 20))
}

func UpdateRow(c *gin.Context) {
	var updateData models.RowTypes
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "字段不正确", "data": gin.H{}, "code": config.ErrorCode})
		return
	}
	hostseditor.Editor.PrintByIndex(updateData.Row - 1)
	item, err := hostseditor.Editor.UpdateLineByIndex(updateData.Row-1, updateData.Hosts, updateData.Ip, updateData.Comments)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": gin.H{}, "code": config.ErrorCode})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "修改成功", "data": item.ToMap(), "code": config.SuccessCode})
}

func DeleteRow(c *gin.Context) {
	var row models.DeleteRowType
	if err := c.BindJSON(&row); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "未设置删除ID!", "data": gin.H{}, "code": config.ErrorCode})
		return
	}

	if item, err := hostseditor.Editor.DeleteLineByIndex(row.Row - 1); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": gin.H{}, "code": config.ErrorCode})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "删除成功!", "data": item.ToMap(), "code": config.SuccessCode})
	}
}

func AddRow(c *gin.Context) {
	var info models.AddRowType
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "参数不全，请填写所有参数!", "data": gin.H{}, "code": config.ErrorCode})
		return
	}

	if item, err := hostseditor.Editor.AddNewRow(info.Hosts, info.Ip, info.Comments); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error(), "data": gin.H{}, "code": config.ErrorCode})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "新增成功!", "data": item.ToMap(), "code": config.SuccessCode})
	}
}
