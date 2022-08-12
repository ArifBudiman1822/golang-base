package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Sidoarjo",
		Province: "Jawa TImur",
		Country:  "",
	}

	ChangeAddressToIndonesia(&alamat)
	fmt.Println(alamat)
}
