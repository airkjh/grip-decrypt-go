package main

import "fmt"

func main() {
	// 빈 인터페이스 : 함수의 매개변수, 리턴 값, 구조체의 필드로 사용 가능
	// 빈 인터페이스 -> 동적 타입 (Java 의 Object 클래스 느낌)
	var a interface{}
	printContents(a)
	
	a = 7.5		// a의 타입이 interface{} 이므로 아무 타입의 값도 할당할 수 있다
	printContents(a)

	a = true
	printContents(a)
}

func printContents(v interface{}) {
	// %T -> 인자의 타입
	fmt.Printf("Type: (%T)\n", v)
	fmt.Printf("Value: (%v)\n", v)
}
