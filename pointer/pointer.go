package main

import "fmt"

type Address struct {
	Kec, Kab, Province string
}

func main() {

	var address1 Address = Address{"Candi", "Sidoarjo", "Jawa Timur"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.Kab = "Mojokerto"

	*address2 = Address{"Ciledug", "Tangerang", "Banten"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.Kec = "Tandes"
	address4.Kab = "Surabaya"
	address4.Province = "Jawa Timur"

	fmt.Println(address4)

}
