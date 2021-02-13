package main

import "fmt"

type test interface{}

func main() {
	/*
		인터페이스
		객체의 동작을 표현, 골격
		동작에 대한 방법만 표시
		추상화 제공
		인터페이스의 메소드를 구현한 타입은 인터페이스로 사용 가능하다
		덕 타이핑 : Go 언어의 독창적인 특성
	*/

	/*
	type 인터페이스명 interface {
		메서드1() 반환값
		메서드2()
		메서드3(인자) 반환값
	}
	 */
	var t test
	fmt.Println(t)			// nil
}
