package main

import (
	"fmt"
)

type account struct {
	number   string
	balance  float32
	interest float32
}

func (account account) Calculate() float32 {
	return account.balance * account.interest
}

func main() {
	// 다양한 선언 방법
	// &struct, struct : &struct 포인터를 받아오기 역참조를 또 하기 때문에 속도가 좀 느리다

	// 방법 1
	var kim *account = new(account) // 포인터 형 변수 주의
	kim.number = "1234"
	kim.balance = 1000000
	kim.interest = 0.015
	fmt.Printf("%#v\n", kim) // 구조체의 모든 멤버 출력

	// 방법 2
	hong := &account{number: "123", balance: 6546464, interest: 0.0456}
	fmt.Printf("%#v\n", hong)

	// 방법 3
	lee := new(account)
	lee.number = "12345"
	lee.balance = 1030000
	lee.interest = 0.01467
	fmt.Printf("%#v\n", lee)

	fmt.Println(int(kim.Calculate()))
	fmt.Println(int(hong.Calculate()))
	fmt.Println(int(lee.Calculate()))
}
