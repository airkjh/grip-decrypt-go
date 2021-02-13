package lib

import "fmt"

func init() {
	fmt.Println("Lib.go initialized")
}

func Greeting(name string) string {
	return "Hello " + name
}

func CheckNum(c int32) bool {
	return c > 10
}
