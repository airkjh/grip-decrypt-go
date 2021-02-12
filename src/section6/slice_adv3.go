package main

import "fmt"

func main() {
	// slice 복사
	// copy(대상, 원본)
	// 반드시 make로 공간을 할당 후 복사해야 한다
	// []int{} 형태의 slice는 복사할 수 없다
	// 복사된 slice의 값을 변경해도 원본에 영향을 주지 않는다
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println(slice2) // [1,2,3,4,5] <- 용량이 5로 지정되었으므로 용량만큼만 복사
	fmt.Println(slice3) // [] <- 용량이 지정되지 않았으므로 복사되지 않음

	slice4 := make([]int, len(slice1)) // 용량을 충분히 확보해야 한다
	copy(slice4, slice1)

	slice4[0] = 100
	fmt.Println(slice4)

	// 주의 사항
	c := []int{1, 2, 3, 4, 5}
	d := c[0:3] // 추출해서 가져오는 것은 참조형이 된다

	d[1] = 7

	fmt.Println(c)
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 0~4번 인덱스까지 가져오되 추출된 새 slice의 capacity는 7로 설정

	fmt.Println("e", e)
	fmt.Printf("f, len = %d, cap = %d, value = %v\n", len(f), cap(f), f)
}
