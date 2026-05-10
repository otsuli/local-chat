package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		err := ServerError{ErrorCode: 67, ErrorMsg: "Something went wrong while opening socket"}
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept() // Wait for someone to connect to port, when they connect Accept() hands a conn representing the connection.
		if err != nil {
			err := ServerError{ErrorCode: 41, ErrorMsg: "Something went wrong while connecting to port"}
			fmt.Println(err)
		}
		go handleConnection(conn) // Spawn a goroutine. Runs handleConnection(conn) in the background concurrently.
	}
}
