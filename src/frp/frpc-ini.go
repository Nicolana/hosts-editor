package frp

import (
	"github.com/Nicolana/hosts-editor/src/server/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

type ServerConfig struct {
	server_addr string `ini:"server_addr"`
	server_port int    `ini:"server_port"`
	token       string
}

type frpcIni struct {
	cfg          *ini.File
	sections     []*ini.Section
	serverConfig *ServerConfig
}

const ServerConfigName = "common"

var FrpcIni = new(frpcIni)

func (f *frpcIni) LoadFile() error {
	cfg, err := ini.Load(IniConfigPath)
	if err != nil {
		return err
	}
	f.cfg = cfg
	f.sections = cfg.Sections()
	return nil
}

func (f *frpcIni) Save() error {
	return f.cfg.SaveTo(IniConfigPath)
}

func (f *frpcIni) GetCommon() *ServerConfig {
	common := &ServerConfig{}
	commonSec := f.cfg.Section(ServerConfigName)
	common.server_port, _ = strconv.Atoi(commonSec.Key("server_port").Value())
	common.server_addr = commonSec.Key("server_addr").Value()
	common.token = commonSec.Key("token").Value()
	return common
}

// UpdateCommon 更新Common的数据
func (f *frpcIni) UpdateCommon(key, value string) (*ServerConfig, error) {
	f.cfg.Section(ServerConfigName).Key(key).SetValue(value)
	err := f.Save()
	if err != nil {
		return &ServerConfig{}, err
	}
	return f.GetCommon(), nil
}

// GetList 获取端口映射的Section列表
func (f *frpcIni) GetList(page, size int, search string) gin.H {
	var sections []gin.H
	for _, value := range f.cfg.Sections() {
		var name = value.Name()
		if name == ServerConfigName || name == "DEFAULT" {
			// 如果是服务器配置参数，则不返回
			continue
		}
		if search != "" {
			// 搜索
			if !strings.Contains(name, search) {
				continue
			}
		}
		var items = make(gin.H)
		for _, key := range value.Keys() {
			items[key.Name()] = key.Value()
		}
		items["name"] = name // Section name
		sections = append(sections, items)
	}

	var total = len(sections)
	if total > 1 {
		var left = (page - 1) * size
		var right = page * size
		if right >= len(sections) {
			right = len(sections)
		}
		sections = sections[left:right]
	}
	return gin.H{
		"page":  page,
		"size":  size,
		"total": total,
		"data":  sections,
	}
}

// UpdateSecByName 根据Sec名称更新该Sec中的映射数据
func (f *frpcIni) UpdateSecByName(item models.FrpSectionType) (*ServerConfig, error) {
	sec := f.cfg.Section(item.Name)
	err := f.Save()
	if err != nil {
		return &ServerConfig{}, err
	}
	return f.GetCommon(), nil
}
