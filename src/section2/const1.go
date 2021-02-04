package main

import "fmt"

func main() {
	// const
	const a int = 1
	fmt.Println("a = ", a)

	const b, c string = "hello", "world"
	const (
		d string = "Discovery"
		e string = "go"
	)

	fmt.Println(a, b, c, d, e)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("Default = ", DEFAULT, "Silver = ", SILVER, "Gold = ", GOLD, "Platinum = ", PLATINUM)
}
