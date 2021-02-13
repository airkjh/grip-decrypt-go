package main

import (
	"fmt"
	"strconv"
)

func helloGoLang() {
	fmt.Println("Hello Go Lang!")
}

func main() {
	/*
		func <함수명>(매개변수) 리턴타입 {}
		func <함수명>() 리턴타입 {}
		func <함수명>(매개변수) {}
		func <함수명>() {}
		다른 언어와 달리 반환값이 다수 가능
	*/
	helloGoLang()
	sayOne("Hello World")
	fmt.Println("sum = ", sum(10, 20))
	fmt.Println("sum op = ", calc(10, 20, sum))

	_min, _max := minMax(10, 20)
	fmt.Println("minmax = ", _min, _max)
	fmt.Println("toString = ", toString(100))

	a := 10
	fmt.Println("a = ", a)
	updateValue(&a, 20)
	fmt.Println("a = ", a)

	fmt.Println("(10 + 20) * 2 = ", nMultiplier(10, 20)(2))

	fmt.Println(multiply(10, 20))
}

// named multiple return value
func multiply(x, y int) (r1, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func updateValue(a *int, value int) {
	*a = value
}

// main 함수 밑에 선언해도 된다
func sayOne(m string) {
	fmt.Println(m)
}

func sum(a int, b int) int {
	return a + b
}

// 함수를 매개 변수로 받는 타입
func calc(a, b int, operator func(int, int) int) int {
	return operator(a, b)
}

func minMax(a, b int) (int, int) {
	return min(a, b), max(a, b)
}

func min(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max(values ...int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func toString(a int) string {
	return strconv.Itoa(a)
}

func nMultiplier(a, b int) func(c int) int {
	c := a + b
	return func(val int) int {
		return c * val
	}
}
