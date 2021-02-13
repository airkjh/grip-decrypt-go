package main

import "fmt"

func main() {
	// 슬라이츠 추가 및 병합
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	//s1 = append(s1, 6, 7, 8)
	fmt.Println(s1)

	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      // slice를 사용할 때는 ... 사용
	s3 = append(s2, s3[0:3]...) // 0부터 3개 -> 0,1,2로 slice

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("len = %d, cap = %d, value = %v\n", len(s4), cap(s4), s4)
	}
}
