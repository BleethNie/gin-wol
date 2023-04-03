package controller

import (
	"fmt"
	"github.com/BleethNie/gin-wol/utils"
	"github.com/BleethNie/gin-wol/utils/r"
	"github.com/gin-gonic/gin"
)

type DeviceController struct{}

// 查询所有设备的信息
func (*DeviceController) QueryDeviceList(c *gin.Context) {
	subnet := utils.GetIntParam(c, "subnet")
	fmt.Println(subnet)
	list := utils.GetDeviceInfoList("")
	fmt.Println(list)
	r.SuccessData(c, list)

}

// 更新设备信息,mac是主键
func (*DeviceController) UpdateDeviceInfo(c *gin.Context) {

}

// 发送wol,进行网络唤醒
func (*DeviceController) Wol(c *gin.Context) {

}

// 是否在线
func (*DeviceController) getDeviceStatus(c *gin.Context) {

}
