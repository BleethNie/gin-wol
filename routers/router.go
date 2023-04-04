package routes

import (
	"github.com/BleethNie/gin-wol/config"
	"github.com/BleethNie/gin-wol/routers/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 后台管理页面的接口路由
func BackRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	// 静态文件服务
	r.Use(static.Serve("/", static.LocalFile("dist/admin", true)))
	r.StaticFS("/dist", http.Dir("./dist"))

	//r.Static("/public", "./public")
	//r.StaticFS("/dir", http.Dir("./public")) // 将 public 目录内的文件列举展示

	r.Use(middleware.Cors()) // 跨域中间件

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		//获取局域网内所有设备信息
		base.GET("/queryDeviceList", deviceController.QueryDeviceList)
		base.GET("/queryDbDeviceList", deviceController.QueryDbDeviceList)
		//更新设备信息
		base.POST("/updateDeviceInfo", deviceController.UpdateDeviceInfo)
		//唤醒
		base.POST("/wol", deviceController.Wol)
	}
	return r
}
