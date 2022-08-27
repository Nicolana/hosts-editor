package models

import "gorm.io/gorm"

type FrpSectionType struct {
	gorm.Model
	Name                 string `gorm:"primary_key" json:"name" binding:"required"`
	NetType              string `gorm:"type:text" json:"type" ini:"type"`
	LocalIp              string `gorm:"type:text" json:"local_ip" ini:"local_ip"`
	LocalPort            int    `json:"local_port" ini:"local_port"`
	RemotePort           int    `json:"remote_port" ini:"remote_port"`
	Sk                   string `gorm:"type:text" json:"sk" ini:"sk"`
	BandwithLimit        string `gorm:"type:text" json:"bandwidth_limit" ini:"bandwidth_limit"`
	UseEncryption        bool   `json:"use_encryption" ini:"use_encryption"`
	UseCompression       bool   `json:"use_compression" ini:"use_compression"`
	Group                string `gorm:"type:text" json:"group" ini:"group"`
	GroupKey             string `gorm:"type:text" json:"group_key" ini:"group_key"`
	HealthCheckType      string `gorm:"type:text" json:"health_check_type" ini:"health_check_type"`
	HealthCheckTimeoutS  string `gorm:"type:text" json:"health_check_timeout_s" ini:"health_check_timeout_s"`
	HealthCheckMaxFailed string `gorm:"type:text" json:"health_check_max_failed" ini:"health_check_max_failed"`
	HealthCheckIntervalS string `gorm:"type:text" json:"health_check_interval_s" ini:"health_check_interval_s"`
	HttpUser             string `gorm:"type:text" json:"http_user" ini:"http_user"`
	HttpPwd              string `gorm:"type:text" json:"http_pwd" ini:"http_pwd"`
	Subdomain            string `gorm:"type:text" json:"subdomain" ini:"subdomain"`
	CustomDomains        string `gorm:"type:text" json:"custom_domains" ini:"custom_domains"`
	Locations            string `gorm:"type:text" json:"locations" ini:"locations"`
	HostHeaderRewrite    string `gorm:"type:text" json:"host_header_rewrite" ini:"host_header_rewrite"`
	Plugin               string `gorm:"type:text" json:"plugin" ini:"plugin"`
	PluginUnixPath       string `gorm:"type:text" json:"plugin_unix_path" ini:"plugin_unix_path"`
	BindAddr             string `gorm:"type:text" json:"bind_addr" ini:"bind_addr"`
	BindPort             string `gorm:"type:text" json:"bind_port" ini:"bind_port"`
	ServerName           string `gorm:"type:text" json:"server_name" ini:"server_name"`
	Role                 string `gorm:"type:text" json:"role" ini:"role"`
}

type FrpDelSectionType struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}

type FrpServerCommon struct {
	Name       string
	ServerAddr string `form:"server_addr" json:"server_addr" xml:"server_addr"  binding:"required"`
	ServerPort int    `form:"server_port" json:"server_port" xml:"server_port" binding:"required"`
	Token      string `form:"token" json:"token" xml:"token" binding:"required"`
}

type FrpServer struct {
	gorm.Model
	FrpServerCommon
}
