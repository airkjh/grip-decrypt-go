package main

import "fmt"

func main() {
	var (
		name      string = "Hello"
		height    int32
		width     float32
		isRunning bool = true
	)

	fmt.Println(name, ",", height, ",", width, ",", isRunning)
}
