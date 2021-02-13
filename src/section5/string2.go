package main

import (
	"fmt"
	"strings"
)

func main() {
	// 추출 - python string slice와 유사하게
	str1 := "Golang"
	fmt.Println(str1[0:2], str1[0]) // str[0] => ascii code
	fmt.Println(str1[2:])
	fmt.Println(str1[:4])

	// 한글의 slicing..?
	str2 := "안녕하세요"
	chars := []rune(str2)
	fmt.Println(chars[0:2])

	// 비교
	str3 := "GoLang"
	str4 := "World"

	fmt.Println("1) is same?", str3 == str4)
	fmt.Println("2) is same?", str3 != str4)
	fmt.Println("3) is same?", str3 == str4)

	// ascii 코드에 대한 사전식 비교
	// G의 ascii code보다 W의 ascii 코드가 숫자가 더 높으므로 false
	fmt.Println("4) is same?", str3 >= str4)

	// 조합
	str5 := "Hello"
	str6 := "World"
	fmt.Println(str5 + " " + str6)

	strs := []string{"Hello", "World"}
	fmt.Println(strings.Join(strs, " "))

	var strs2 []string
	strs2 = append(strs2, "Hello")
	strs2 = append(strs2, "World")
	fmt.Println(strings.Join(strs2, " "))

}
