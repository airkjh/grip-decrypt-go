package main

import "fmt"

func main() {
	var char1 rune = 50556 		// unicode
	var char2 rune = 0142574
	var char3 rune = 0xC57C

	// %c -> 한 글자
	// %d -> 10진수 숫자
	// %o -> 8진수
	// %x -> 16진수

	fmt.Printf("%c %c %c\n", char1, char2, char3)
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	fmt.Printf("%d %o %x\n", char1, char2, char3)
}
