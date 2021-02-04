package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("ex1) i = ", i)
	}

	// 에러 발생 패턴 - {} 위치
	//for i := 0; i < 5; i++
	//{
	//
	//}

	// 에러 발생 패턴2 - 1줄짜리 본문이라서 {}를 생략한 경우
	//for i := 0; i < 5; i++
	//	fmt.Println(i)

	// 무한 루프 패턴
	//for {
	//	fmt.Println("Hello Go Lang")
	//}
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3) i = ", index, ", name = ", name)
	}
	for _, name := range loc {
		fmt.Println("ex3) name = ", name)
	}
}
