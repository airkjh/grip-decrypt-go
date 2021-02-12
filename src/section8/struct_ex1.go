package main

import "fmt"

type Account2 struct {
	number   string  "계좌번호"
	balance  float64 "잔액"
	interest float64 "이자율"
}

// 생성자 패턴
// 리턴 값이 포인터가 아니면 값 복사이기 때문에 포인터 타입으로 리턴해야 한다
func newAccount2(number string, balance float64, interest float64) *Account2 {
	return &Account2{number, balance, interest}
}

func main() {
	// 생성자와 유사한 패턴
	kim := Account2{number: "1234", balance: 1000, interest: 0.012}

	//
	lee := new(Account2)
	lee.number = "1234"
	lee.balance = 1000
	lee.interest = 0.012

	fmt.Println(kim, lee)

	park := newAccount2("ABCD", 123123, 0.12312)
	fmt.Println(park)

}
