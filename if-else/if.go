package main

import (
	"fmt"
)

func main() {
	var studentScore = 110

	if studentScore > 100 || studentScore < 0 {
		fmt.Println("Unknown")
	} else if studentScore >= 80 {
		fmt.Println("Nilai A")
	} else if studentScore >= 65 {
		fmt.Println("Nilai B")
	} else if studentScore >= 50 {
		fmt.Println("Nilai C")
	} else if studentScore >= 0 {
		fmt.Println("Nilai D")
	}
}
