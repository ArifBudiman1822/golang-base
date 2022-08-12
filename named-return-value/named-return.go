package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "Muhammad"
	middlename = "Arif"
	lastname = "Budiman"
	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
