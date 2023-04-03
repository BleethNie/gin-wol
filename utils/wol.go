package utils

import (
	"fmt"
	"net"
)

func wol(mac string) {
	// 将 MAC 地址转换为 6 字节的字节数组
	dstMAC, err := net.ParseMAC(mac)
	if err != nil {
		fmt.Printf("无法解析 MAC 地址：%v\n", err)
		return
	}

	// 构建 WOL 魔术包
	magicPacket := make([]byte, 102)
	for i := 0; i < 6; i++ {
		magicPacket[i] = 0xff
	}
	for i := 1; i < 17; i++ {
		for j := 0; j < 6; j++ {
			magicPacket[i*6+j] = dstMAC[j]
		}
	}

	// 发送 WOL 魔术包广播
	conn, err := net.Dial("udp", "255.255.255.255:9")
	if err != nil {
		fmt.Printf("无法发送 WOL 魔术包：%v\n", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write(magicPacket)
	if err != nil {
		fmt.Printf("无法发送 WOL 魔术包：%v\n", err)
		return
	}
	fmt.Println("WOL 魔术包已发送。")
}
