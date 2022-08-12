package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	registerUser("Anjeng", func(name string) bool {
		return name == "Anjeng"
	})

	registerUser("Arif", func(name string) bool {
		return name == "Anjeng"
	})
}
