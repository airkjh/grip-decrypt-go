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

func (d dog) bite() {
	fmt.Printf("Dog %s bites!\n", d.name)
}

func (d dog) bark() {
	fmt.Printf("Dog %s barks!\n", d.name)
}

func (d dog) run() {
	fmt.Printf("Dog %s runs!\n", d.name)
}

func (d cat) bite() {
	fmt.Printf("Cat %s bites!\n", d.name)
}

func (d cat) bark() {
	fmt.Printf("Cat %s barks!\n", d.name)
}

func (d cat) run() {
	fmt.Printf("Cat %s runs!\n", d.name)
}

type Animal interface {
	bite()
	bark()
	run()
}

func act(animal Animal) {
	animal.bite()
	animal.bark()
	animal.run()
}

func main() {
	// 덕 타이핑 -> 메서드로만 타입을 판단한다
	dog := dog{"D", 10.0}
	cat := cat{"C", 10.0}

	act(dog)
	act(cat)
}
