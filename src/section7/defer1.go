package main

import "fmt"

func main() {
	/*
		finally 같은 기능 (try...catch..finally)
		Defer 함수 -> 지연 실행
		Defer를 호출한 함수가 종료되기 직전에 호출된다
		리소스를 반환하거나 파일 닫기, Mutex 잠금 해제
	*/
	exF1()
}

func exF1() {
	fmt.Println("f1 : start")
	defer exF2()
	fmt.Println("f1 : End")
}

func exF2() {
	fmt.Println("f2 : called")
}
