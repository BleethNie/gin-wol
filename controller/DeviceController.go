package controller

import (
	"fmt"
	"github.com/BleethNie/gin-wol/dao"
	model "github.com/BleethNie/gin-wol/model/entity"
	"github.com/BleethNie/gin-wol/utils"
	"github.com/BleethNie/gin-wol/utils/r"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeviceController struct{}

// 查询所有设备的信息
func (*DeviceController) QueryDeviceList(c *gin.Context) {
	subnet := c.Query("subnet")
	fmt.Println(subnet)
	deviceList := utils.GetDeviceInfoList(subnet)
	r.SuccessData(c, deviceList)
}

// 查询所有数据库里的数据
func (*DeviceController) QueryDbDeviceList(c *gin.Context) {
	entityList := dao.List([]model.DeviceEntity{}, "*", "", "")
	r.SuccessData(c, entityList)

}

// 更新设备信息
func (*DeviceController) UpdateDeviceInfo(c *gin.Context) {
	var device model.DeviceEntity
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := dao.Count(model.DeviceEntity{}, "mac = ?", device.Mac)
	if count != 0 {
		dao.Updates(&device, "mac = ?", device.Mac)
	} else {
		dao.Create(&device)
	}
	r.SuccessMessage(c, "更新成功")
}

// 发送wol,进行网络唤醒
func (*DeviceController) Wol(c *gin.Context) {
	mac := c.Query("mac")
	hostname := c.Query("hostname")
	nickname := c.Query("nickname")
	if mac != "" {
		utils.WakeOnLAN(mac)
	}
	if hostname != "" {
		deviceEntity := dao.GetOne(model.DeviceEntity{}, "host_name = ?", hostname)
		if deviceEntity.Mac != "" {
			utils.WakeOnLAN(mac)
		}
	}
	if nickname != "" {
		deviceEntity := dao.GetOne(model.DeviceEntity{}, "nick_name = ?", nickname)
		if deviceEntity.Mac != "" {
			utils.WakeOnLAN(mac)
		}
	}
	r.SuccessMessage(c, "发送WOL成功")
}

// 是否在线
func (*DeviceController) ClearDbAndSave(c *gin.Context) {
	dao.Delete(model.DeviceEntity{}, "ip like ?", "%.%")

	subnet := c.Query("subnet")
	deviceList := utils.GetDeviceInfoList(subnet)
	for _, value := range deviceList {
		dao.Create(&value)
	}
	r.SuccessData(c, deviceList)
}

// 是否在线
func (*DeviceController) getDeviceStatus(c *gin.Context) {

}
