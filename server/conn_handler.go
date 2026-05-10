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
	buffer := make([]byte, 1024) // 1kb slice.

}
