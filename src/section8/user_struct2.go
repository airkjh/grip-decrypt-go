package main

import "fmt"

type cnt int

type totalCost func(int, int) int

func main() {
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본타입, 함수.. 재정의 가능
	a := cnt(15) // cnt가 함수인것처럼 보이기 때문에 이렇게는 잘 사용하지 않는다

	var b cnt = 15

	fmt.Println(a)
	fmt.Println(b)

	// 호출 불가! cnt 타입도 int 타입이지만 타입은 정확하게 사용해야 한다.
	// 정 쓰고 싶으면 int() 이용해서 형변환 해야한다
	//testConverT(a)
	testConverD(a)

	var orderPrice totalCost
	orderPrice = func(cnt int, price int) int {
		return cnt * price
	}

	describe(10, 100, orderPrice)
}

func testConverT(i int) {
	fmt.Println("(Default Type)", i)
}

func testConverD(i cnt) {
	fmt.Println("(Custom Type)", i)
}

func describe(cnt int, price int, fn totalCost) {
	fmt.Printf("cnt = %d, price = %d, totalCost = %d\n", cnt, price, fn(cnt, price))
}
