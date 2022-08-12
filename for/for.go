package main

import "fmt"

func main() {
	//for counter <= 10 {
	//	fmt.Println("Perulangan Ke", counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Berulang Ke", counter)

	}

	//name := "Muhammad"

	//for i := 0; i < len(name); i++ {
	//	fmt.Println("Kata ke", i, string(name[i]))
	//}

	slice := []string{"Muhammad", "Arif", "Budiman"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	handphone := make(map[string]string)
	handphone["merk"] = "Xiomi"
	handphone["type"] = "POCO F3"

	for key, value := range handphone {
		fmt.Println(key, "=", value)
	}

}
