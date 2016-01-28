package utils

import (
	"net"
	"strings"
)

func GetLocalIp() ([]string, error) {
	info, _ := net.InterfaceAddrs()
	var res []string
	for _, addr := range info {
		res = append(res, strings.Split(addr.String(), "/")[0])
	}
	return res, nil
}

func GetExclusiveLocalIp() ([]string, error) {
	info, _ := net.InterfaceAddrs()
	var res []string
	for _, addr := range info {
		if !strings.EqualFold(strings.Split(addr.String(), "/")[0], "0.0.0.0") {
			res = append(res, strings.Split(addr.String(), "/")[0])
		}
	}
	return res, nil
}
