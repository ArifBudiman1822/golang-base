package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) sayMr() {
	man.Name = "Mr " + man.Name
}

func main() {
	var name Man
	name.Name = "Arif"

	name.sayMr()
	fmt.Println(name.Name)
}
