package utils

import (
	"fmt"
	"net"
	"strings"
)

func WakeOnLAN(macAddr string) error {
	if strings.Contains(macAddr, "-") {
		macAddr = strings.ReplaceAll(macAddr, "-", ":") // 需要将 MAC 地址中的 " - " 替换为 " : "
	}
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		return err
	}

	hwAddr, err := net.ParseMAC(macAddr)
	if err != nil {
		return err
	}

	udpConn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer udpConn.Close()

	buf := []byte{}
	for i := 0; i < 6; i++ {
		buf = append(buf, 0xff)
	}
	for i := 0; i < 16; i++ {
		buf = append(buf, hwAddr...)
	}

	_, err = udpConn.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	macAddr := "00-11-22-33-44-55"
	macAddr = strings.ReplaceAll(macAddr, "-", ":") // 需要将 MAC 地址中的 " - " 替换为 " : "
	err := WakeOnLAN(macAddr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("WOL packet sent successfully to", macAddr)
	}
}
