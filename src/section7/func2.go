package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// 가변인자, 재귀함수, 익명함수 예제
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumAll())

	fmt.Println("10 factorial = ", factorial(10))
	fmt.Println("2 factorial = ", factorial(2))

	a := []int{5, 6, 7, 8, 9, 10}
	m := sumAll(a...)
	fmt.Println(m)

	// 함수의 슬라이스
	funcs := []func(values ...int) int{sumAll, multiplyAll}
	inputs := []int{1, 2, 3, 4, 5}
	for _, _func := range funcs {
		fmt.Println(GetFunctionName(_func), _func(inputs...))
	}

	/*
		익명함수 (Anonymous Function)
		선언과 동시에 함수를 실행하고 싶을 때
	*/
	func(a int) int {
		fmt.Println(a)
		return a
	}(10)
}

//GetFunctionName 는 Refection으로 인자로 넘어온 함수의 이름을 리턴
//외부에서 접근 가능한 함수는 이렇게 주석이 있어야 하고 반드시 함수 이름으로 시작하는 주석이어야 한다
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// 가변 인자
func sumAll(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func multiplyAll(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	m := 1
	for _, val := range values {
		m *= val
	}
	return m
}

// 재귀 함수
func factorial(a int64) int64 {
	if a == 0 || a == 1 {
		return a
	}
	return a * factorial(a-1)
}
