package main

import (
	"fmt"
	"time"
)

func main() {
	// 예제1: 동기, 버퍼 미사용
	ch := make(chan bool)
	cnt := 6
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			// 채널에 데이터를 보내면 수신할 때까지 대기
			fmt.Println("Go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<- ch
		fmt.Println("Main: ", i)
	}
}
