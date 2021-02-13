package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, " start", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>", i)
	}
	fmt.Println(name, " end", time.Now())
}

func main() {
	go exe("T1")
	fmt.Println("Main Routine start")
	go exe("T2")
	go exe("T3")
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine end")
}