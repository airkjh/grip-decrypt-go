package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 버퍼 사용, 비동기
	runtime.GOMAXPROCS(1)

	// 버퍼 크기 : 2
	// 버퍼 발신 -> 가득차면 대기, 수신 -> 비어있으면 대기
	bufferSize := 2
	ch := make(chan bool, bufferSize)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			// 버퍼가 찰 때까지 채널에 발신
			ch <- true
			fmt.Println("Go :", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		// 버퍼가 다 빌 때까지 채널을 수신
		<-ch
		fmt.Println("Main: ", i)
	}

}
