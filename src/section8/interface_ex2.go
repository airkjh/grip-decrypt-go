package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type assertion
	// 빈 인터페이스 타입을 런타임시에 연산 등을 위해서 실제 타입으로 변환해야 할 경우가 생김
	// 인터페이스.(타입) -> 형 변환
	var a interface{} = 15

	b := a
	c := a.(int)
	//d := a.(float64)		// interface conversion error, 원래 타입으로만 캐스팅 가능

	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	//fmt.Println(d)

	if v, ok := a.(int); ok {
		fmt.Println(v, ok)
	}

	// typeCastable이 false이므로 아래 if는 실행되지 않는다
	if v, typeCastable := a.(float64); typeCastable {
		fmt.Println(v, ok)
	}
}
