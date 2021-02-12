package main

import "fmt"

type dog struct {
	name   string
	weight float32
}

type cat struct {
	name   string
	weight float32
}

func (d dog) run() {
	fmt.Printf("Dog %s runs!\n", d.name)
}

func (d cat) run() {
	fmt.Printf("Cat %s runs!\n", d.name)
}

type Animal interface {
	run()
}

// 익명 인터페이스 예제
func act(animal interface {run()}) {
	animal.run()
}

func main() {
	dog := dog{"D", 10.0}
	cat := cat{"C", 10.0}

	act(dog)
	act(cat)
}
