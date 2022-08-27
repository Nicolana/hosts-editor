package global

import (
	"github.com/Nicolana/hosts-editor/src/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() *gorm.DB {
	dsn := "root:123456@tcp(192.168.10.103:3306)/switch?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// 迁移schema, 执行Migrate会生成对应的表
	db.AutoMigrate(&models.FrpSectionType{})
	db.AutoMigrate(&models.FrpServer{})

	// 将DB导出
	// Db = db
	return Db
}
