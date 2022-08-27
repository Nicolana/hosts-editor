package global

import "gorm.io/gorm"

type config struct {
	Db *gorm.DB
}

var Config = new(config)

func SystemInit() {
	Config.Db = InitDatabase()
}
