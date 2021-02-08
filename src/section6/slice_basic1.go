package main

import "fmt"

/*
	슬라이스 기초
	배열과 유사하지만 길이가 가변,
	레퍼런스 타입
	길이 & 용량
	배열처럼 선언하거나 make(자료형, 길이, 용량) 함수 사용
 */

func main() {
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1,2,3,4,5}

	fmt.Printf("%-5T %p %d %d %v\n", slice1, &slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %p %d %d %v\n", slice2, &slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %p %d %d %v\n", slice3, &slice3, len(slice3), cap(slice3), slice3)

	slice4 := make([]int, 5, 10)
	fmt.Printf("%-5T %p %d %d %v\n", slice4, &slice4, len(slice4), cap(slice4), slice4)

	//slice4 := make([]int, 5) --> 이렇게 생성하면 length 5, capacity 5로 선언

	var slice5 []int		// nil, 길이와 용량이 0
	fmt.Println(slice5)

	if slice5 == nil {
		fmt.Println("slice5 is nil")
	}
}
