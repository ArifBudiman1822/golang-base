package main

import "fmt"

type s string

func getFullName() (s, s) {
	return "Arif", "Budiman"
}

func main() {

	firstname, lastname := getFullName()
	fmt.Println(firstname)
	fmt.Println(lastname)

}
