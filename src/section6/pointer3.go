package main

import "fmt"

func rptc(n *int) {
	// n의 레퍼런스 복사
	*n = 77
}

func vptc(n int) {
	// n의 값 복사
	n = 77
}

func main() {
	// 크기가 큰 배열인 경우 값 복사시 시스템에 부담 -> 포인터로 전달
	a := 10
	b := 10

	vptc(a)
	fmt.Println(a)

	rptc(&b)
	fmt.Println(b)
}
