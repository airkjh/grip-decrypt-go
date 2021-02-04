package main

import "fmt"

func main() {
	// 짧은 선언
	// 전역 사용 불가, 함수 내에서만 사용
	// 선언 후 재할당하면 에러
	shortVar := 3

	fmt.Println(shortVar)
}
