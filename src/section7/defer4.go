package main

import "fmt"

func main() {
	a()
}

func a() {
	// defer의 범위는 end 함수까지만... start 함수는 바로 실행됨
	// 안티 패턴
	defer end(start("b"))
	fmt.Println("in a")
}

func start(s string) string {
	fmt.Println(s)
	return s
}

func end(e string) {
	fmt.Println("end = ", e)
}
