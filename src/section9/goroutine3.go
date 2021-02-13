package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>", r, i)
	}
	fmt.Println(name, " end ", time.Now())
}

func main() {
	// 멀티 코어 (멀티 CPU) 최대한 활용
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	// numCPU = runtime.GOMAXPROCS(0) 으로 가져올 수 있다
	fmt.Println("Current System CPU", numCPU)

	fmt.Println("Main Routine start", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine end", time.Now())
}