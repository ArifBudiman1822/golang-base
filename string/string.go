package main

import "fmt"

func main() {
	// len() == menghitung panjang string
	fmt.Println("Nama Saya")
	fmt.Println("Muhammad")
	fmt.Println(len("Muhammad"))
	fmt.Println("Arif")
	fmt.Println(len("Arif"))

	var name string = "Muhammad"
	fmt.Println(len(name))
	fmt.Println(name[0]) // byte
	fmt.Println(string(name[0]))
}
