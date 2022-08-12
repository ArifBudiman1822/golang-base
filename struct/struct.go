package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var arif Customer
	arif.Name = "Muhammad Arif Budiman"
	arif.Address = "Sidoarjo"
	arif.Age = 21

	fmt.Println(arif)
	fmt.Println(arif.Name)
	fmt.Println(arif.Address)
	fmt.Println(arif.Age)

	sahril := Customer{
		Name:    "Sahril Mahendra",
		Address: "Kost D Three - Tangerang",
		Age:     24,
	}

	fmt.Println(sahril)

	pimen := Customer{"Firmansyah Ainul Yakin", "Malang", 20}

	fmt.Println(pimen)
}
