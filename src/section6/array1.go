package main

import "fmt"

/*
	(배열 : 길이 고정) vs (슬라이스 : 길이 가변)
	(배열 : 값 타입 -> 값을 복사 전달, 비교 연산자 사용가능) vs (슬라이스: 참조 타입 -> 참조 값 전달, 비교 연산자 사용 가능)

	cap() -> 배열, 슬라이스의 용량
	len() -> 배열, 슬라이스 개수
*/

func main() {
	// 배열
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // capacity와 갯수가 일치해야 한다
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [...]int{1, 2, 3, 4, 5} // 크기를 확신할 수 없을 때, 배열 크기가 자동 증가

	fmt.Println("arr1=", arr1)
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)
	fmt.Println("arr4=", arr4)
	fmt.Println("arr5=", arr5)

	// 2차원 배열
	arr6 := [5][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	// %-5T : Type 정보
	fmt.Printf("%-5T, %d, %v\n", arr6, len(arr6), arr6)

	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Printf("%-5T, %d, %v\n", arr9, len(arr9), arr9)

}
