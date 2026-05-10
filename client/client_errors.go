package main

type clientError struct {
	ErrorMsg  string
	ErrorCode int
	User      user
}
