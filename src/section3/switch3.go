package main

import "fmt"

func main() {
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("one of 2,4,6")
	case 1, 3, 5:
		if a > 20 {
			fmt.Println("Too big")
		} else {
			fmt.Println("one of 1,3,5")
		}
	default:
		fmt.Println("not matched")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("java")
		fallthrough // 조건에 맞지 않아도 다음 case 문으로 진입
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
	case "ruby":
		fmt.Println("ruby")
	}

}
