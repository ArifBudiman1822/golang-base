package main

import "fmt"

func random() interface{} {
	return true
}

func main() {

	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown Type")
	}

}