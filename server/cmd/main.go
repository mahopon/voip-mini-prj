package main

import (
	"github.com/mahopon/voip-mini-prj/internal/server"
	"log"
)

func main() {
	addr := "0.0.0.0"
	host := 12345

	log.Printf("UDP server starting on %s:%d", addr, host)
	server.StartServer(addr, host)
}
