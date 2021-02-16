package main

import "fmt"

func main() {

	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good"
		}
	}()

	val1, ok1 := <- ch
	fmt.Println(val1, ok1)
	val2, ok2 := <- ch
	fmt.Println(val2, ok2)
	val3, ok3 := <- ch
	fmt.Println(val3, ok3)

	close(ch)

	val4, ok4 := <- ch
	fmt.Println(val4, ok4)

}
