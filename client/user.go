package client

import "math/rand/v2"

type user struct {
	name string
	UID  uint8
}

func newUser(name string) user {
	return user{
		name: name,
		UID:  uint8(rand.IntN(129)),
	}
}
