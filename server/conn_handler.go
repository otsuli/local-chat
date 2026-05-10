package main

import "net"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Figure out what I want to do with the connection:
}
