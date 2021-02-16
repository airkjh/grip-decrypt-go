package main

import (
	"fmt"
	"time"
)

func main() {
	// 함수의 매개 변수에 수신 및 발신 전용 채널 지정 가능
	// 전용 채널 설정 후 방향이 다를 때 예외 발생
	// 발신 전용 : chan <- 데이터형
	// 수신 전용 : <- chan
	// 매개 변수를 통해서 전용 채널을 확인할 수 있다
	// 함수의 반환값으로 채널 사용 가능
	c := make(chan int)

	go sendOnly(c, 10) // 발신 전용

	go receiveOnly(c) // 수신 전용

	time.Sleep(3 * time.Second)
}

// 수신 전용 채널 사용
// 이 함수 내에서는 채널 수신만 가능
func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received", i)
	}

	fmt.Println(<-c)
}

// 발신 전용 채널 사용
// 이 함수 내에서는 채널에 데이터 발신만 가능
func sendOnly(c chan<- int, cnt int) {
	for i := 0; i < cnt; i++ {
		c <- i
	}
	c <- 777
}
