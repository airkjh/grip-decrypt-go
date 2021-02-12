package main

import "fmt"

type Account struct {
	number   string
	balance  float32
	interest float32
}

func (account Account) Calculate() float32 {
	return account.balance * account.interest
}

func main() {

	// 클래스, 객체, 클래스 개념 없음
	// 리시버를 통해 메서드를 연결한다
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스 (리시버)
	a := Account{number: "123", balance: 10000000.0, interest: 0.05}
	b := Account{number: "345", balance: 2465489761.0, interest: 0.015}

	fmt.Println(a.Calculate())
	fmt.Println(b.Calculate())

}
