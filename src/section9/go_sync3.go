package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 동기화 사용하지 않은 경우
	// 쓰기 읽기 동작 순서가 일정하지 않아서
	// 잘못된 오류를 반환할 가능성 증가
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write: ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1 :", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2 :", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
