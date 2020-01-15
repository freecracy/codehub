package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	ip4 = flag.Bool("4", false, "ip 4 地址")
)

func main() {
	if *ip4 != true {
		fmt.Println("error")
		os.Exit(0)
	}
	getip4()
}

func getip4() {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		ip, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ip.IP.IsLoopback() {
			continue
		}
		if ip.IP.To4() != nil {
			fmt.Println(ip.IP.String())
		}
	}
}

func init() {
	flag.Parse()
}
