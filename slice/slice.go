package main

import "fmt"

func main() {

	var day = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	var slice1 = day[1:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//day[2] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Diganti"
	//fmt.Println(day)

	var slice2 = day[5:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Arif")
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Arif"
	newSlice[1] = "Budiman"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
