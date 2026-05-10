package main

import "fmt"

type ServerError struct {
	ErrorMsg  string
	ErrorCode int
}

func (e ServerError) Error() string {
	return fmt.Sprintf("error %d: %s", e.ErrorCode, e.ErrorMsg)
}
