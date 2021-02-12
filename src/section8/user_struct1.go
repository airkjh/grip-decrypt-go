package main

import "fmt"

type car struct {
	name  string
	color string
	price int64
	tax   int64
}

// 일반 함수
func Price(c car) int64 {
	return c.price + c.tax
}

// 구조체, 함수 바인딩
func (c car) Price() int64 {
	return c.price + c.tax
}

func main() {
	// Go는 클래스 없음 ( 대신 구조체 )
	// 전형적인 객체 지향 특징을 가지고 있지 않지만
	// 인터페이스를 통한 다형성, 구조체를 이용한 클래스 형태의 코딩 가능
	// 상태, 메소드 분리해서 정의 ( 클래스에 다 몰아두는 구조 아님)
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본타입, 함수.. 재정의 가능
	bmw := car{name: "520D", color: "white", price: 50000000, tax: 500000}
	benz := car{name: "220D", color: "white", price: 60000000, tax: 600000}

	fmt.Printf("%v, %p\n", bmw, &bmw)
	fmt.Printf("%v, %p\n", benz, &benz)

	// 이렇게는 잘 사용하지 않는다
	fmt.Println(Price(bmw))
	fmt.Println(Price(benz))

	// 이런 방식으로 주로 사용한다
	fmt.Println(bmw.Price())
	fmt.Println(benz.Price())

	fmt.Println(&bmw == &benz)

}
