package models

//	type RowTypes struct {
//		Row      int    `form:"row" json:"row" xml:"row" binding:"required"`
//		Hosts    string `form:"hosts" json:"hosts" xml:"hosts"  binding:"required"`
//		Ip       string `form:"ip" json:"ip" xml:"ip" binding:"required"`
//		Comments string `form:"comments" json:"comments"`
//	}
type FrpSectionType struct {
	Name       string `form:"name" json:"name" xml:"name" binding:"required"`
	NetType    string `form:"type" json:"type" xml:"type"  binding:"required"`
	LocalIp    string `form:"local_ip" json:"local_ip" xml:"local_ip" binding:"required"`
	LocalPort  int    `form:"local_port" json:"local_port" xml:"local_port" binding:"required"`
	RemotePort int    `form:"remote_port" json:"remote_port" xml:"remote_port" binding:"required"`
}

type FrpDelSectionType struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}
