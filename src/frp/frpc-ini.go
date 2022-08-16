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

func (f *frpcIni) GetCommon() gin.H {
	common := &ServerConfig{}
	commonSec := f.cfg.Section(ServerConfigName)
	common.server_port, _ = strconv.Atoi(commonSec.Key("server_port").Value())
	common.server_addr = commonSec.Key("server_addr").Value()
	common.token = commonSec.Key("token").Value()
	return gin.H{
		"server_addr": common.server_addr,
		"server_port": common.server_port,
		"token":       common.token,
	}
}

// UpdateCommon 更新Common的数据
func (f *frpcIni) UpdateCommon(item models.FrpCommonSection) (gin.H, error) {
	// 如果当前服务器数据不存在，则创建该数据Section
	if !f.cfg.HasSection(ServerConfigName) {
		if _, err := f.cfg.NewSection(ServerConfigName); err != nil {
			return gin.H{}, err
		}
	}
	sec := f.cfg.Section(ServerConfigName)
	sec.Key("server_addr").SetValue(item.ServerAddr)
	sec.Key("server_port").SetValue(strconv.Itoa(item.ServerPort))
	sec.Key("token").SetValue(item.Token)
	if err := f.Save(); err != nil {
		return gin.H{}, err
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

func (f *frpcIni) ToMap(item ini.Section) (gin.H, error) {
	return gin.H{
		"type":        item.Key("type").Value(),
		"local_ip":    item.Key("local_ip").Value(),
		"local_port":  item.Key("local_port").Value(),
		"remote_port": item.Key("remote_port").Value(),
	}, nil
}

// UpdateSecByName 根据Sec名称更新该Sec中的映射数据
func (f *frpcIni) UpdateSecByName(item models.FrpSectionType) (gin.H, error) {
	sec := f.cfg.Section(item.Name)
	sec.Key("type").SetValue(item.NetType)
	sec.Key("local_ip").SetValue(item.LocalIp)
	sec.Key("local_port").SetValue(strconv.Itoa(item.LocalPort))
	sec.Key("remote_port").SetValue(strconv.Itoa(item.RemotePort))
	err := f.Save()
	if err != nil {
		return gin.H{}, err
	}
	return f.ToMap(*sec)
}

// AddNewSec 新增Sec数据
func (f *frpcIni) AddNewSec(item models.FrpSectionType) (gin.H, error) {
	sec, err := f.cfg.NewSection(item.Name)
	if err != nil {
		return gin.H{}, err
	}
	sec.Key("type").SetValue(item.NetType)
	sec.Key("local_ip").SetValue(item.LocalIp)
	sec.Key("local_port").SetValue(strconv.Itoa(item.LocalPort))
	sec.Key("remote_port").SetValue(strconv.Itoa(item.RemotePort))
	err = f.Save()
	if err != nil {
		return gin.H{}, err
	}
	return f.ToMap(*sec)
}

// DelSec 根据Name删除一条Section
func (f *frpcIni) DelSec(name string) (gin.H, error) {
	sec := f.cfg.Section(name)
	f.cfg.DeleteSection(name)
	err := f.Save()
	if err != nil {
		return gin.H{}, err
	}
	return f.ToMap(*sec)
}
