package main

import (
	"fmt"
	"runtime"
	"time"
	"sync"
)

func main() {
	// 읽기 전용 락, 쓰기 전용 락을 사용하여 공유 데이터 보호
	// RWMutex -> 쓰기 락, 쓰기 시도 중 다른 곳에서 읽지도 쓰지도 못한다
	// RMutex -> 읽는 동안 쓰기 방지
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock()
			data += 1
			fmt.Println("Write: ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read1 :", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()
			fmt.Println("Read2 :", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
