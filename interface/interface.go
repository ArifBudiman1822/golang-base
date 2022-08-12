package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Handphone struct {
	Name string
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (handphone Handphone) GetName() string {
	return handphone.Name
}

func main() {
	var arif Person
	arif.Name = "Arif"
	sayHello(arif)

	var merk Handphone
	merk.Name = "POCO F3"
	sayHello(merk)
}
