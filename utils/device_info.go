package utils

import (
	"container/list"
	"fmt"
	model "github.com/BleethNie/gin-wol/model/entity"
	"github.com/axgle/mahonia"
	"log"
	"net"
	"os/exec"
	"regexp"
	"runtime"
	"sync"
)

func GetDeviceInfoList(subnet string) *list.List {
	ips, err := getActiveIPs(subnet)
	if err != nil {
		log.Fatal(err)
	}
	deviceList := list.New()

	var wg sync.WaitGroup
	for i := 0; i < len(ips); i++ {
		wg.Add(1)
		ip := ips[i]
		go func(ip string) {
			var deviceEntity model.DeviceEntity
			hostname, err1 := getHostname(ip)
			if err1 == nil {
				mac := getMacAddr(ip)
				fmt.Printf("%s ==> %s ==> %s \n", ip, hostname, mac)

				deviceEntity.Ip = ip
				deviceEntity.Mac = mac
				deviceEntity.HostName = hostname

				deviceList.PushBack(deviceEntity)
			}
			wg.Done()
		}(ip)
	}
	wg.Wait()

	return deviceList
}

func getMacAddr(ip string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("arp", "-a", ip)
	} else {
		cmd = exec.Command("arp", "-n", ip)
	}

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	outString := string(out)
	if runtime.GOOS == "windows" {
		outString = gBKToUTF8(outString)
	}
	mac := extractMACAddress(outString)
	return mac
}

func getActiveIPs(subnet string) ([]string, error) {
	var ips []string
	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return ips, err
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips, nil
}

func getHostname(ip string) (string, error) {
	names, err := net.LookupAddr(ip)
	if err != nil {
		return ip, err
	}

	if len(names) > 0 {
		return names[0], nil
	}

	return ip, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// 将 GBK 编码的字符串转换为 UTF-8 编码的字符串
func gBKToUTF8(gbk string) string {
	srcCoder := mahonia.NewDecoder("gbk")
	res := srcCoder.ConvertString(gbk)
	return res
}

func extractMACAddress(str string) string {
	// 定义正则表达式
	re := regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)
	match := re.FindString(str)
	return match
}
