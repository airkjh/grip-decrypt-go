package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 Wait -> 잠들게 하기, Signal, Broadcast 올 때까지 즉시 작업을 중단함
	// Signal -> 고루틴 하나만 깨우기, Broadcast -> 모둔 잠든 고루틴 깨우기
	runtime.GOMAXPROCS(runtime.NumCPU())

	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	// Buffer size가 5이므로 비동기처럼 실행됨
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("GoRoutine waiting", n)
			condition.Wait()		// 고루틴 대기
			fmt.Println("Waiting End", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		//<-c
		fmt.Println("Received", <-c)
	}

	// 고루틴 하나씩 깨우기
	// 깨우기 전에 락 걸어줘야 함
	//for i := 0; i < 5; i++ {
	//	mutex.Lock()
	//	fmt.Println("Wake go routine(signal)", i)
	//	condition.Signal()
	//	mutex.Unlock()
	//}

	mutex.Lock()
	condition.Broadcast()
	mutex.Unlock()


	time.Sleep(2 * time.Second)
}
