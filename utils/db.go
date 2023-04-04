package utils

import (
	"github.com/BleethNie/gin-wol/config"
	model "github.com/BleethNie/gin-wol/model/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitSqlite() *gorm.DB {
	path := config.Cfg.Db.Path
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
// 只支持创建表、增加表中没有的字段和索引
// 为了保护数据，并不支持改变已有的字段类型或删除未被使用的字段
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.DeviceEntity{},
	)

	if err != nil {
		log.Println("gorm 自动迁移失败: ", err)
	} else {
		log.Println("gorm 自动迁移成功")
	}
}
