package main

import "fmt"

func main() {
	// 익명 구조체
	car1 := struct{ name, color string }{"320D", "White"}

	fmt.Printf("%#v\n", car1)

	cars := []struct{ name, color string }{
		{"Sonata", "White"},
		{"520D", "Black"},
		{"320D", "Grey"},
	}

	for _, car := range cars {
		fmt.Printf("%#v\n", car)
	}
}
