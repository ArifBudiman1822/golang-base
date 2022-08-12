package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "Nama saya", customer.Name)
}

func (b Customer) sayHei() {
	fmt.Println("Hei Pelanggan!!!", b.Name)
}

func main() {
	var arif Customer
	arif.Name = "Arif"
	arif.Address = "Sidoarjo"
	arif.Age = 21

	arif.sayHello("Joko")

	arif.sayHei()
}
