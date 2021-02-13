package main

import "fmt"

func main() {
	// 배열 값 복사 증명
	arr1 := [5]int{1, 2, 3, 4, 5}
	v1 := arr1[0]
	fmt.Println(v1)

	v1 = 7
	fmt.Println(v1)

	arr2 := arr1

	// %p => pointer
	// 메모리 주소가 다르므로 값이 복사된 것이다
	fmt.Printf("%p %v\n", &arr1[0], arr1[0])
	fmt.Printf("%p %v\n", &arr2[0], arr2[0])
}
