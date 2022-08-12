package main

import "fmt"

func main() {
	const phi float64 = 3.14
	const name string = "Muhammad Arif Budiman"
	const benar bool = true

	fmt.Println(phi)
	fmt.Println(name)
	fmt.Println(benar)

	const (
		firstname = "Arif"
		lastname  = "Budiman"
	)

	fmt.Println(firstname, lastname)
}
