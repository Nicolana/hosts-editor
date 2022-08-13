package models

type RowTypes struct {
	Row      int    `form:"row" json:"row" xml:"row" binding:"required"`
	Hosts    string `form:"hosts" json:"hosts" xml:"hosts"  binding:"required"`
	Ip       string `form:"ip" json:"ip" xml:"ip" binding:"required"`
	Comments string `form:"comments" json:"comments"`
}

type DeleteRowType struct {
	Row int `form:"row" json:"row" xml:"row" binding:"required"`
}

type AddRowType struct {
	Hosts    string `form:"hosts" json:"hosts" xml:"hosts"  binding:"required"`
	Ip       string `form:"ip" json:"ip" xml:"ip" binding:"required"`
	Comments string `form:"comments" json:"comments"`
}
