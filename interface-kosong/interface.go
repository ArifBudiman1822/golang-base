package main

import (
	"fmt"
)

func tesAja(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups Salah"
	}
}

func main() {

	printinteger := tesAja(1)
	fmt.Println(printinteger)
	printbool := tesAja(2)
	fmt.Println(printbool)
	printstring := tesAja(3)
	fmt.Println(printstring)
}
