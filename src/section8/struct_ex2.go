package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee     // 구조체 임베디드 패턴, Go 상속 X, 메서드 상속을 위한 패턴
	specialBonus float64
}

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	emp1 := Employee{"kim", 2000000, 150000}
	emp2 := Employee{"park", 1500000, 200000}
	ex1 := Executives{
		Employee{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println(emp1, int(emp1.Calculate()))
	fmt.Println(emp2, int(emp2.Calculate()))

	// Employee 구조체를 통해서 메서드 호출
	fmt.Println(ex1, int(ex1.Employee.Calculate()+ex1.specialBonus))
	fmt.Println(ex1, int(ex1.Calculate()))

	// 구조체 임베디드 메서드 오버라이드 패턴
}
