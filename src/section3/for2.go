package main

import "fmt"

func main() {

	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}

	fmt.Println("sum1 = ", sum1)

	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++ // 후치 연산은 반환값이 없음
	}
	fmt.Println("sum2 = ", sum2)

	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("sum3 = ", sum3)

	for i, j := 0, 0; i < 10; i, j = i+1, j+10 {
		fmt.Println("ex4) i = ", i, ", j = ", j)
	}
}
