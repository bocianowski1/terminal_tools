package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Get the local address from the connection
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.IP.String())
}
