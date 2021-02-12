package main

import "fmt"

func main() {
	// 빈 인터페이스는 어떠한 자료형도 전달받을 수 있으므로 타입 체크를 통해 형 변환 후 사용 가능
	checkType(true)
	checkType(1)
	checkType(22.54)
	checkType(nil)
	checkType("Hello GoLang")
}

func checkType(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("This is bool")
	case int, int8, int16, int32, int64:
		fmt.Println("This is integer")
	case float32, float64:
		fmt.Println("This is float")
	case string:
		fmt.Println("This is string")
	case nil:
		fmt.Println("This is nil")
	default:
		fmt.Println("This is unknown")
	}
}