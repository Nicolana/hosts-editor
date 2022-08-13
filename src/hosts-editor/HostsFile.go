package hostseditor

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type HostsItem struct {
	index        int // 标记行数
	hosts        string
	ip           string
	hasComments  bool   // 标记当前行有没有注释
	comments     string // 存放注释
	originString string // 存放原始字符串
}

type HostsFile struct {
	Lines []HostsItem
}

func (item HostsItem) ToString() string {
	var res = ""
	if item.ip != "" && item.hosts != "" {
		res = item.ip + " " + item.hosts
	}
	if item.hasComments {
		comments := item.comments
		if !strings.HasPrefix(strings.TrimSpace(comments), "#") {
			// 防止没有 # 的情况被加入到行内去导致文件损坏
			comments = "#" + comments
		}
		if item.ip != "" && item.hosts != "" {
			res += " " + comments + GetReturnSymbol()
		} else {
			res = comments
		}
	} else {
		res += GetReturnSymbol()
	}
	return res
}

// 将有效地Hosts转成Map 切片
func (item HostsItem) ToMap() gin.H {
	return gin.H{
		"index":        item.index,
		"hosts":        item.hosts,
		"ip":           item.ip,
		"hasComments":  item.hasComments,
		"comments":     item.comments,
		"originString": item.originString,
	}
}

// 将列表转换成列表对象
func (hosts HostsFile) ToMapArray() []gin.H {
	var list []gin.H
	for _, value := range hosts.Lines {
		list = append(list, value.ToMap())
	}
	return list
}

// 将Hosts列表转换成列表对象
func (hosts HostsFile) HostsToMapArray() []gin.H {
	var list []gin.H
	for _, value := range hosts.Lines {
		if !value.hasComments || (value.ip != "" && value.hosts != "") {
			list = append(list, value.ToMap())
		}
	}
	return list
}
