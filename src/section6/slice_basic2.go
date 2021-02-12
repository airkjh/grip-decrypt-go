package main

import "fmt"

func main() {
	// slice는 진짜 참조 타입인가?
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 // 복사
	arr2[0] = 7

	fmt.Println(arr1) // [1,2,3]
	fmt.Println(arr2) // [7,2,3]

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1 // 참조 연결
	slice2[0] = 7

	fmt.Println(slice1) // [7,2,3]
	fmt.Println(slice2) // [7,2,3]

	// 슬라이스 예외 상황
	slice3 := make([]int, 50, 100)
	fmt.Println(slice3[4]) // int.default로 초기화 -> 0 출력
	// fmt.Println(slice3[50])		// out of range! 초기화는 length만큼만, capacity만큼 아님

}
