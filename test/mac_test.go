package test

import (
	"bufio"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/wonderivan/logger"
)

func TestGetLANDevIP(t *testing.T) {
	//这里的设备名称可通过ifconfig查找
	GetLANDevIP("eth0")
}

func GetLANDevIP(devName string) ([]string, error) {
	//清除对应设备的arp缓存

	time.Sleep(time.Second * time.Duration(5))
	//目前openwrt上似乎不是安装的arp命令，而是软链接的查看该文件，所以arp命令实际上应该是如下内容
	cmd := exec.Command("arp", "-a")
	output, err := cmd.Output()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Debug(string(output))

	var res []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
		if strings.Contains(scanner.Text(), devName) {
			tmp := strings.Split(scanner.Text(), " ")
			res = append(res, tmp[0])
		}
	}
	return res, nil
}
