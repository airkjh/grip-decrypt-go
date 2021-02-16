package main

import "fmt"

func main() {
	// 채널 또한 함수의 리턴 타입으로 사용 가능
	c := sum(100)
	fmt.Println("ex1 : ", <-c)
}

// 수신 전용 채널 리턴
func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 0; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}