package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	nanoTime := time.Now().UnixNano()

	fmt.Println("[Unix Seconds]", time.Now().Unix())
	fmt.Println("[Unix Milli]", nanoTime/int64(time.Millisecond))
	fmt.Println("[Unix Nano] seed value = ", nanoTime)

	fmt.Println(time.Millisecond, int64(time.Millisecond))

	rand.Seed(nanoTime)

	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i = ", i, " 50이상 100미만")
	case i > 25 && i < 50:
		fmt.Println("i = ", i, " 25이상 50미만")
	default:
		fmt.Println("i = ", i, " 기본값")
	}

}
