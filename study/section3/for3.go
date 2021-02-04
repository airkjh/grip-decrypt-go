package main

import "fmt"

func main() {
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1 // Loop label이 없으면 i == 2, j == 4일때 j loop를 빠져나가고 i 루프를 계속 진행하게 됨
			}
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
Loop2:
	// Loop 레이블 아래는 바로 for loop가 와야한다
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3)", i, j)
		}
	}
}
