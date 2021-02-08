package main

import "fmt"

func main() {
	/*
		주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
		포인터 *
		nil로 초기화 ( 0 != nil )
	 */

	var a *int			// 포인터형, 주소값만 저장 가능, nil로 초기화
	var b *int = new(int)			// 랜덤한 주소로 초기화

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)		// 실제 주소로 가보면 0으로 자료형의 기본값으로 초기화되어 있음

	i := 7
	fmt.Println(i, &i)

	a = &i
	fmt.Println(*a, i)		// 7 7

	i = 10
	fmt.Println(*a, i)		// 10 10

	*a = 123
	fmt.Println(*a, i)		// 123 123



}
