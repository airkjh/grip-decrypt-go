package main

import "fmt"

func main() {
	// defer function은 stack 구조(후입선출)로 호출된다
	stack()
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
