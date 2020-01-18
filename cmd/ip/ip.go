package main

import (
	"flag"
	"net"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

const VersionDefaultUsage string = "output Version"

var (
	Version string
	v       bool
)

var (
	ip4 *bool = flag.Bool("4", false, "ip 4 地址")
)

func main() {

	flag.BoolVar(&v, "v", false, VersionDefaultUsage)
	flag.BoolVar(&v, "version", false, VersionDefaultUsage)

	flag.Parse()
	log.SetHandler(cli.Default)

	if v {
		log.Info(Version)
		return
	}

	if *ip4 {
		getip42()
	}

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
			log.Info(ip.IP.String())
		}
	}
}

func getip42() {
	conn, err := net.Dial("udp", "github.com:80")
	if err != nil {
		return
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	log.Info(ip)
}
