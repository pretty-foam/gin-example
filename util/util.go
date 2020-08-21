package util 

import (
	"os"
	"net"
	"fmt"
)


//GetIP 获取IP
func GetIP() string {

	var ip string

	addrs ,err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _ , address := range addrs{
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip 
}