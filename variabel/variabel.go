package main

import "fmt"

func main() {
	var name string = "Muhammad"
	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(string(name[0]))

	var age int = 21
	fmt.Println(age)

	var married bool = false
	fmt.Println(married)

	var (
		firstname  = "Muhammad"
		middlename = "Arif"
		lastname   = "Budiman"
	)

	fmt.Println(firstname, middlename, lastname)

}
