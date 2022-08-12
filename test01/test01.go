package main

import "fmt"

func kaliSemua(barisangka ...int) int {
	total := 1

	for _, value := range barisangka {
		total *= value
	}
	return total
}

func main() {
	result := kaliSemua(3, 3, 3)
	fmt.Println(result)
}
