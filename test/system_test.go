package test

import (
	"fmt"
	"github.com/BleethNie/gin-wol/utils"
	"runtime"
	"testing"
)

func TestSysType(t *testing.T) {

	sysType := runtime.GOOS
	if sysType == "linux" {
		fmt.Println("nice")
	}
	if sysType == "windows" {
		fmt.Println("windows")
	}

}

func TestCmd(t *testing.T) {
	out, err := utils.RunArp()
	if err != nil {
		fmt.Println("error")
	}
	utils.ParseMac(out)
}
