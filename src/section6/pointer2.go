package main

import "fmt"

func main() {
	i := 7
	p := &i

	fmt.Println(i, p)
	fmt.Println(&i, p)
	fmt.Println(i, *p)

	*p++
	fmt.Println(*p, i)		// 8 8


}
