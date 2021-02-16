package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel Select 구문
	// 채널에 값이 수신되면 해당 case문을 실행
	// 일회성 구문이므로, For 안에서 수행
	// default 구문 처리 주의
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			// 이런 패턴에서는 default 사용하지 말 것!
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2: ", str)
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
