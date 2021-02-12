package main

import "fmt"

type dog struct {
	name   string
	weight float32
}

type cat struct {
	name   string
	weight float32
}

// Any type allowed
func printValue(s interface {}) {
	fmt.Println(s)
}

func main() {
	// 인터페이스 활용
	// 함수내에서 어떤 타입이라도 유연하게 매개변수로 받을 수 있다 -> 모든 타입 지정 가능
	dog := dog {"D", 10.0}
	cat := cat {"C", 9.0}

	printValue(dog)
	printValue(cat)

}
