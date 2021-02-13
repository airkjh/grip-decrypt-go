package main

import "fmt"

type Dog struct {
	name   string
	weight float32
}

func (d Dog) bite() {
	fmt.Printf("%s bites!\n", d.name)
}

type Behavior interface {
	bite()
}

func main() {
	// 인터페이스 구현
	dog1 := Dog{"Miso", 10.0}
	var inter1 Behavior
	inter1 = dog1
	inter1.bite()
	fmt.Printf("%#v\n", dog1)

	// 이 패턴이 더 많이 쓰임
	dog2 := Dog {"Mimi", 9.0}
	inter2 := Behavior(dog2)
	inter2.bite()

	animals := []Behavior{dog1, dog2}
	for _, animal := range animals {
		animal.bite()
	}
}
