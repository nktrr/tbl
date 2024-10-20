package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:4000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		inputBytes := make([]byte, 1024)
		conn.ReadFromUDP(inputBytes)
		fmt.Println(string(inputBytes))
	}

}
