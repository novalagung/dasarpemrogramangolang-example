package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var ip net.IP
	flag.TextVar(&ip, "ip", net.IPv4(127, 0, 0, 1), "IP address")

	flag.Parse()
	fmt.Println("ip:", ip)
}
