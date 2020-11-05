package database

import (
	"fmt"
)

var Users = make(map[string]User)

func init() {
	addUser("Billy")
	addUser("Chris")
	addUser("Tim")
	addUser("Kyle")
}

func addUser(name string) {
	Users[name] = User{
		Name: name,
		ID:   fmt.Sprintf("user:%s", name),
	}
}

func AllUsersList() []*User {
	reply := make([]*User, len(Users))
	var i int
	for _, u := range Users {
		reply[i] = &u
		i++
	}

	return reply
}
