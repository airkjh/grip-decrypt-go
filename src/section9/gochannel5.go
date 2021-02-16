package main

import "fmt"

func main() {
	// Close : 채널 닫기
	// 주의 : 닫힌 채널에 갑 전송시 패닉(예외) 발생
	// Range : 채널 안에서 값을 꺼낸다.
	// 채널 닫아야 반복문 종료 -> 채널이 열려 있고 값 전송하지 않으면 계속 대기
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 아래에서 계속 채널 수신을 대기 하므로 채널을 명시적으로 닫지 않으면 프로그램이 종료되지 않는다
	}()

	// 채널값을 순호하며 계속 수신
	// 채널이 닫힐 때까지 채널 값을 수신
	for i := range ch {
		fmt.Println("ex1", i)
	}
}
