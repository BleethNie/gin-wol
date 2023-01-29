package utils

import (
	"fmt"
	"github.com/BleethNie/gin-wol/config"
	"github.com/spf13/viper"
	"log"
)

func InitViper() {
	// 根据命令行读取配置文件路径
	var configPath string
	log.Println(fmt.Sprintf("加载配置文件路径: config/config.toml"))
	configPath = fmt.Sprintf("config/config.toml")

	v := viper.New()
	v.SetConfigFile(configPath)

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}

	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}

	log.Println("配置文件内容加载成功")
}
