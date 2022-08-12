package main

import (
	"errors"
	"fmt"
)

func tesPembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		hasil := nilai / pembagi
		return hasil, nil
	}
}

func main() {
	hasil, err := tesPembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
