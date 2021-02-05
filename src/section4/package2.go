package main

import (
	"fmt"
	myLib "section4/lib"
)

func main() {
	greetingMsg := myLib.Greeting("airkjh")
	fmt.Println(greetingMsg)

	myLib.GetOut("Dumb ass")
}
