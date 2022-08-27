package models

//	type RowTypes struct {
//		Row      int    `form:"row" json:"row" xml:"row" binding:"required"`
//		Hosts    string `form:"hosts" json:"hosts" xml:"hosts"  binding:"required"`
//		Ip       string `form:"ip" json:"ip" xml:"ip" binding:"required"`
//		Comments string `form:"comments" json:"comments"`
//	}
type FrpSectionType struct {
	Name                 string `json:"name" binding:"required"`
	NetType              string `json:"type" ini:"type"`
	LocalIp              string `json:"local_ip" ini:"local_ip"`
	LocalPort            int    `json:"local_port" ini:"local_port"`
	RemotePort           int    `json:"remote_port" ini:"remote_port"`
	Sk                   string `json:"sk" ini:"sk"`
	BandwithLimit        string `json:"bandwidth_limit" ini:"bandwidth_limit"`
	UseEncryption        bool   `json:"use_encryption" ini:"use_encryption"`
	UseCompression       bool   `json:"use_compression" ini:"use_compression"`
	Group                string `json:"group" ini:"group"`
	GroupKey             string `json:"group_key" ini:"group_key"`
	HealthCheckType      string `json:"health_check_type" ini:"health_check_type"`
	HealthCheckTimeoutS  string `json:"health_check_timeout_s" ini:"health_check_timeout_s"`
	HealthCheckMaxFailed string `json:"health_check_max_failed" ini:"health_check_max_failed"`
	HealthCheckIntervalS string `json:"health_check_interval_s" ini:"health_check_interval_s"`
	HttpUser             string `json:"http_user" ini:"http_user"`
	HttpPwd              string `json:"http_pwd" ini:"http_pwd"`
	Subdomain            string `json:"subdomain" ini:"subdomain"`
	CustomDomains        string `json:"custom_domains" ini:"custom_domains"`
	Locations            string `json:"locations" ini:"locations"`
	HostHeaderRewrite    string `json:"host_header_rewrite" ini:"host_header_rewrite"`
	Plugin               string `json:"plugin" ini:"plugin"`
	PluginUnixPath       string `json:"plugin_unix_path" ini:"plugin_unix_path"`
	BindAddr             string `json:"bind_addr" ini:"bind_addr"`
	BindPort             int    `json:"bind_port" ini:"bind_port"`
	ServerName           int    `json:"server_name" ini:"server_name"`
	Role                 int    `json:"role" ini:"role"`
}

type FrpDelSectionType struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}

type FrpCommonSection struct {
	ServerAddr string `form:"server_addr" json:"server_addr" xml:"server_addr"  binding:"required"`
	ServerPort int    `form:"server_port" json:"server_port" xml:"server_port" binding:"required"`
	Token      string `form:"token" json:"token" xml:"token" binding:"required"`
}
