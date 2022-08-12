package main

import "fmt"

func main() {
	// && jika salah satu false == false
	// || jika salah satu true == true

	var scoreujian = 80
	var absensi = 70
	var lulusujian = scoreujian <= 80
	var lulusabsensi = absensi <= 70
	var lulus bool = lulusujian && lulusabsensi

	fmt.Println(lulus)

}
