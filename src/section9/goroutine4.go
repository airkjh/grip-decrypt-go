package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 클로저 사용 예제
	runtime.GOMAXPROCS(1)

	s := "Goroutine Closure : "

	// Go routine의 클로저는 가장 나중에 실행
	// 즉 반복문을 넘어가는 순간 실행된다
	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
	}
	time.Sleep(3 * time.Second)
}
