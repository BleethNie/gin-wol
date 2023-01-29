package utils

import (
	"github.com/BleethNie/gin-wol/config"
	"github.com/BleethNie/gin-wol/dao"
	"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"gorm.io/gorm"
	"log"
)

func InitSQLite(migrateFlag string) *gorm.DB {
	sqlLiteCfg := config.Cfg.SQLite

	db, err := gorm.Open(sqlite.Open(sqlLiteCfg.Name), &gorm.Config{})

	if err != nil {
		log.Fatal("sqlite 连接失败, 请检查参数")
	}

	log.Println("sqlite 连接成功")

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	if migrateFlag == "false" {
		autoMigrate(db)
	}
	return db
}

// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
// 只支持创建表、增加表中没有的字段和索引
// 为了保护数据，并不支持改变已有的字段类型或删除未被使用的字段
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dao.DeviceDao{},
	)

	if err != nil {
		log.Println("gorm 自动迁移失败: ", err)
	} else {
		log.Println("gorm 自动迁移成功")
	}
}
