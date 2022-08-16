package main

import (
	"fmt"
	"math"
)

func main() {
	// membulatkan sesuai yang terdekat
	fmt.Println(math.Round(1.6))
	fmt.Println(math.Round(1.3), "\n")

	//membulatkan ke bawah
	fmt.Println(math.Floor(1.6))
	fmt.Println(math.Floor(1.3), "\n")

	//membulatkan ke atas
	fmt.Println(math.Ceil(1.6))
	fmt.Println(math.Ceil(1.4), "\n")

	// memilih nilai terbesar
	fmt.Println(math.Min(10, 20))
	fmt.Println(math.Max(10, 20))

}
