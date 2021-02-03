package main

import "fmt"

func main() {
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음")
	}

	switch d := "go"; d + "lang" {
	case "golang":
		fmt.Println("Go Lang")
	case "javalang":
		fmt.Println("Java Lang")
	default:
		fmt.Println("None")
	}

	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i == j")
	case i > j:
		fmt.Println("i > j")
	}
}
