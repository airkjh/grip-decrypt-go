package main

import "fmt"

type car1 struct {
	color string
	name  string
}

func main() {
	c1 := car1{"red", "220D"}
	c2 := new(car1)
	c2.color, c2.name = "white", "Sonata"
	c3 := &car1{}
	c4 := &car1{"black", "520D"}

	fmt.Printf("%#v\n", c1)
	fmt.Printf("%#v\n", c2)
	fmt.Printf("%#v\n", c3)
	fmt.Printf("%#v\n", c4)
}
