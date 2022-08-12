package main

import "fmt"

func sayHelloWithName(firstname, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	firstname := "Arif"
	sayHelloWithName(firstname, "Budiman")
	sayHelloWithName("Budi", "Kurap")
}
