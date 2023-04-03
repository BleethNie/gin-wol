package utils

import (
	"errors"
	model "github.com/BleethNie/gin-wol/model/entity"
	"github.com/mingrammer/commonregex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"runtime"
	"strings"
)

type Buffer struct {
	Data  []byte
	start int
}

func (b *Buffer) PrependBytes(n int) []byte {
	length := cap(b.Data) + n
	newData := make([]byte, length)
	copy(newData, b.Data)
	b.start = cap(b.Data)
	b.Data = newData
	return b.Data[b.start:]
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

// 反转字符串
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func ParseMac(data string) []model.DeviceEntity {
	var deviceEntityList []model.DeviceEntity

	macList := strings.Split(data, "\r")
	for i := 0; i < len(macList); i++ {
		macs := commonregex.MACAddresses(macList[i])
		ips := commonregex.IPv4s(macList[i])

		if len(macs) == 0 || len(ips) == 0 {
			continue
		}
		if macs[0] == "ff-ff-ff-ff-ff-ff" || strings.Contains(ips[0], "255.255.255") {
			continue
		}
		if strings.Contains(ips[0], "224.0.0") {
			continue
		}
		deviceEntity := model.DeviceEntity{Mac: macs[0], Ip: ips[0], NickName: "", HostName: ""}
		deviceEntityList = append(deviceEntityList, deviceEntity)
	}
	return deviceEntityList
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

func RunArp() (string, error) {
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
