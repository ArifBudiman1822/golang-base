package main

import "fmt"

func main() {

	var customer map[string]string = map[string]string{
		"name":    "Arif",
		"address": "Sidoarjo",
	}

	customer["title"] = "Junior Programmer"

	fmt.Println(customer)
	fmt.Println(customer["name"])
	fmt.Println(customer["address"])

	var laptop = make(map[string]string)
	laptop["merk"] = "ACER"
	laptop["ram"] = "4GB"
	laptop["kondisi"] = "Busuk"
	fmt.Println(laptop)
	fmt.Println(len(laptop))

	delete(laptop, "kondisi")

	fmt.Println(laptop)
	fmt.Println(len(laptop))

}
