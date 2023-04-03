package test

import (
	"fmt"
	utils "github.com/BleethNie/gin-wol/utils"
	"testing"
)

func TestGetHostname(t *testing.T) {
	list := utils.GetDeviceInfoList("192.168.2.0/24")
	fmt.Println(list)

}
