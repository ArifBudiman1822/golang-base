package helper

import "fmt"

var application = "Belajar Golang"
var Version = "1.0.0"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
