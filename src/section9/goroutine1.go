package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		비동기적 함수 루틴 실행(매우 적은 용량 차이) -> 채널을 이용한 통신
		공유 메모리 사용시에 정확한 동기화 코딩 필요
	*/
	exe1()
	fmt.Println("Main Routine start", time.Now())

	// go routine 으로 실행
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine end", time.Now())
}

func exe1() {
	fmt.Println("exe1 function start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 function end", time.Now())
}

func exe2() {
	fmt.Println("exe2 function start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 function end", time.Now())
}

func exe3() {
	fmt.Println("exe3 function start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 function end", time.Now())
}
