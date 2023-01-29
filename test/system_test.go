package test

import (
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
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

func runArp() (string, error) {
	sysType := runtime.GOOS
	if sysType == "linux" {
		out, err := exec.Command("arp").Output()
		return string(out), err
	}
	if sysType == "windows" {
		out, err := exec.Command("arp", "-a").Output()
		return ConvertByte2String(out, "GB18030"), err
	}
	return "", errors.New("系统不匹配")
}

func TestCmd(t *testing.T) {
	out, err := runArp()
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(out)
}

func ConvertByte2String(byte []byte, charset string) string {
	var str string
	switch charset {
	case "GB18030":
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
