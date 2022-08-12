package main

import "fmt"

type filtered func(string) string

func sayHelloWithFilter(name string, filter filtered) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Asu" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Arif", spamFilter)
	sayHelloWithFilter("Asu", spamFilter)
}
