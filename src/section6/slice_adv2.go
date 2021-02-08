package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스 추출 및 정렬
	// slice[i:j] -> i부터 j-1까지 추출
	// slice[i:] -> i부터 끝까지 추출
	// slice[:j] -> 처음부터 j-1까지 추출
	// slice[:] -> 처음부터 마지막까지 추출
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1[:])
	fmt.Println(slice1[0:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[0:len(slice1)])

	// 정렬
	slice2 := []int{3,1,6,4,7,8,0}
	slice3 := []string{"B", "D", "F", "A", "C", "E"}

	fmt.Println(sort.IntsAreSorted(slice2))			// 이 배열/슬라이스가 정렬되어 있나요?

	sort.Ints(slice2)		// 정렬해주세요
	fmt.Println(slice2)

	fmt.Println()
	fmt.Println(sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println(slice3)
}
