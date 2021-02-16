package main

import "fmt"

func main() {
	c := receiveOnly(100)
	output := total(c)

	fmt.Println(<-output)
}

func receiveOnly(cnt int) <-chan int {
	sum := 0
	// tot는 함수 내에서 선언했기 때문에 함수 내에서는 발신/수신 가능
	// 리턴하는 순간 수신 전용
	tot := make(chan int)
	go func() {
		for i := 0; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

// 수신 전용 채널 리턴
func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a += i
		}
		tot <- a
	}()

	return tot
}
