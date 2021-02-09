package main

import "fmt"

func main() {
	sayHello("GoLang")
}

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi")
	}()
}
