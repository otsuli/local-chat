package main

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			err := ServerError{ErrorMsg: "Error Closing Connection", ErrorCode: 383}
			fmt.Println(err)
		}
	}()
	buffer := make([]byte, 5*1024) // five 1kb byte slices.
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			err := ServerError{ErrorCode: 21, ErrorMsg: "something went wrong while reading incoming data"}
			fmt.Println(err)
		}
		fmt.Println("Packet received from: ", conn.RemoteAddr().String())
		fmt.Println("Packet contents: ", string(buffer[:n])) // Convert slice to string
	}
}
