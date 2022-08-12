package main

import "fmt"

func main() {
	var namasaya = [3]string{"Muhammad", "Arif", "Budiman"}

	fmt.Println(len(namasaya))
	fmt.Println(namasaya)
	fmt.Println(namasaya[0])
	fmt.Println(namasaya[1])
	fmt.Println(namasaya[2])

	var age [3]int

	age[0] = 21
	age[1] = 9
	age[2] = 45

	fmt.Println(len(age))
	fmt.Println(age)
	fmt.Println(age[0])
	fmt.Println(age[1])
	fmt.Println(age[2])
}
