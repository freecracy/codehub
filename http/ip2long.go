package http

import (
	"errors"
	"math"
	"net"
)

func ip2long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("d")
	}
	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

func long2ip(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "nil", errors.New("err")
	}
	b := make(net.IP, net.IPv4len)
	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)
	return b.String(), nil
}
