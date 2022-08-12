package main

import "fmt"

func main() {

	name := "Rulli"

	switch name {
	case "Arif":
		fmt.Println("Namanya Arif")
	case "Rulli":
		fmt.Println("Namanya Rulli")
	default:
		fmt.Println("G Bener Namanya")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Kepanjangan")
	//case false:
	///	fmt.Println("Wes Bener Uy")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Kepanjangan Nemen")
	case length > 5:
		fmt.Println("Kepanjangan")
	default:
		fmt.Println("Wes Bener Uy")
	}
}
