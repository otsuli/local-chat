package client

import (
	"fmt"
	"net"
)

func init_user() user {
	var name string
	fmt.Print("Type your username: ")
	fmt.Scan(&name)
	fmt.Println("Client Initialized")
}

func main() {
	init_user()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		err := clientError{ErrorCode: 1, ErrorMsg: "Can not connect to client"}
		fmt.Println(err)
	}

}
