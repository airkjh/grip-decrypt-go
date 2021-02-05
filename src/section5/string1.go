package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 string = "c:\\go_study\\src"
	var str2 string = `c:\go_study\src`

	fmt.Println(str1)
	fmt.Println(str2)

	str1 = "안녕하세요"
	fmt.Println(len(str1))		// 한글 1글자는 3바이트

	str2 = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(str2)) // 5글자
	fmt.Println(len([]rune(str2)))		// alternative

	str3 := "Hello"
	fmt.Println(len(str3))

	for i := 0; i < len(str3); i++ {
		// str[i] => ascii code 값 출력
		fmt.Println(string(str3[i]))
	}

	/* 이렇게 하면 str[i]가 1바이트씩 나오니까 정상적으로 출력되지 않는다 */
	fmt.Println("한글 slicing...")
	for i := 0; i < len(str1); i++ {
		fmt.Println(string(str1[1]))
	}

	/* 한글 slicing은 이렇게... */
	fmt.Println("한글 slicing2...")
	chars := []rune(str1)
	for i := 0; i < len(chars); i++ {
		fmt.Println(string(chars[i]))
	}
	/* 아니면 이렇게 */
	for _, char := range str1 {
		fmt.Println(string(char))
	}
}
