package main

import (
	"fmt"
	"runtime"
	"sync"
)

type count struct {
	num int
	mutex sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock()
	c.num += 1
	c.mutex.Unlock()
}

func (c *count) result() int {
	return c.num
}

func main() {

	// Mutex를 이용하여 공유 데이터 보호
	// sync.Mutex 선언 후, Lock, Unlock
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
