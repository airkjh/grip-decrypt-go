package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() int {
	return c.num
}

func main() {

	// Mutex를 사용하지 않는 고 루틴의 데이터 공유
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched()		// 고 루틴이 끝났기 때문에 다른 CPU에 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	fmt.Println(c.result())
}
