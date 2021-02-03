package main

import "fmt"

/*
	iota
 */

func main() {
	const (
		A = iota
		B
		C
	)

	fmt.Println(A, B, C)

	const (
		D = iota * 10
		E
		F
	)

	fmt.Println(D, E, F

	)
}
