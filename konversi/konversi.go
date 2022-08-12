package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	const name string = "Arif"
	var e uint = uint(name[0])
	var stringName string = string(name[0])

	fmt.Println(e)
	fmt.Println(stringName)
}
