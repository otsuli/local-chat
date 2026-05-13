package main

import (
	"fmt"
	"net"
)

func initUser() user {
	var name string
	var client user
	fmt.Print("Type your username: ")
	if _, err := fmt.Scan(&name); err != nil {
		err := clientError{ErrorCode: 34, ErrorMsg: "Error initializing client username"}
		fmt.Println(err)
	}
	fmt.Println("Client Initialized")
	return client
}

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		err := clientError{ErrorCode: 1, ErrorMsg: "Can not connect to client"}
		fmt.Println(err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			err := clientError{ErrorCode: 911, ErrorMsg: "Error closing listener"}
			fmt.Println(err)
		}
	}()

	client := initUser()
	if _, err := fmt.Fprintf(conn, "%s, %d", client.name, client.UID); err != nil {
		err := clientError{ErrorCode: 48, ErrorMsg: "Error sending username and user ID"}
		fmt.Println(err)
	}

	for {
		var message string
		fmt.Print("message> ")
		if _, err := fmt.Scan(&message); err != nil {
			err := clientError{ErrorCode: 983, ErrorMsg: "Error taking message input"}
			fmt.Println(err)
		}
		if _, err := fmt.Fprintf(conn, message); err != nil {
			err := clientError{ErrorCode: 238, ErrorMsg: "Error sending message"}
			fmt.Println(err)
		}
	}
}
