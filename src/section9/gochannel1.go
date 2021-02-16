package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E---> ", time.Now())
	v <- 1 // 1을 채널로 전송
}

func work2(v chan int) {
	fmt.Println("Work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ---> ", time.Now())
	v <- 2 // 2을 채널로 전송
}

func main() {
	/*
		고 루틴간의 상호 정보(데이타) 교환 및 실행 흐름 동기화 위해서 사용 ( 채널 = 동기식 )
		실행 흐름 제어 (동기, 비동기)
		채널을 일반 변수로 선언 후 사용 가능
		데이터 전달 자료형 선언 후 사용 (지덩된 타입만 주고 받을 수 있음)
		interface{} 전달을 통해서 자료형 상관없이 전송 및 수신 가능
		값을 전달 (복사 후: int, bool 등), 포인터 (슬라이스, 맵) 등을 전달시에는 주의!
		멀티 프로세싱 처리에서 교착 상태 (경쟁) 주의!
		<-, -> 기호 사용 (채널 <- 데이터 : 데이터를 채널로 보낸다, 변수 <- 채널 : 데이터 수신)
	*/
	fmt.Println("Main: S ---> ", time.Now())

	v := make(chan int)
	go work1(v)
	go work2(v)

	<-v
	<-v

	fmt.Println("Main: E ---> ", time.Now())
}
