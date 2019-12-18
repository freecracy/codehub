package http

import "flag"

import "fmt"

// Sendmail ...
func Sendmail() {
	fmt.Println("helllo")
}

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host1", "127.0.0.1", "host name")
	flag.IntVar(&port, "port1", 25, "port")
	flag.Parse()
	fmt.Println(host)
	fmt.Println(port)
}
