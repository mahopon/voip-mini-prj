package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func StartServer(host string, port int) {

	// Resolve the string address to a UDP address
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	udpAddr, err := net.ResolveUDPAddr("udp", addr) // Returns UDPAddr, which is taken in by ListenUDP
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 512)
	// Read from UDP listener in endless loop
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		fmt.Print("> ", string(buf[:n]))

		// Write back the message over UPD
		_, err = conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
		if err != nil {
			fmt.Println("Write error:", err)
		}
	}
}
