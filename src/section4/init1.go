package main

import "fmt"

var needToInitialized bool = false

func init() {
	fmt.Println("Init method start!")
	needToInitialized = true
}

func main() {
	fmt.Println("Main method start!")
	fmt.Println("value of needToInitialized=", needToInitialized)
}
