package main

import (
	"fmt"
	"strings"
)

type User struct {
	Id      int
	Name    string
	Email   string
	Address string
}

func main() {
	users := []User{
		{1, "name-1", "mail@com1", "address1"},
		{2, "name-2", "mail@com2", "address2"},
		{3, "name-3", "mail@com3", "address3"},
		{4, "name-4", "mail@com4", "address4"},
		{5, "name-5", "mail@com5", "address5"},
		{6, "name-6", "mail@com6", "address6"},
		{7, "name-7", "mail@com7", "address7"},
	}

	// the filter func
	f := func(u User) bool {
		return u.Id >= 5 && strings.Contains(u.Email, "6")
	}
	filtered := userFilter(users, f)

	fmt.Println(filtered)
}

func userFilter(users []User, f func(User) bool) []User {
	var u []User
	for _, user := range users {
		if f(user) {
			u = append(u, user)
		}
	}
	return u
}
