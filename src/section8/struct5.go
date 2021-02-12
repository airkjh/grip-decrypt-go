package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전측"
}

func main() {
	// 필드 사용
	tags := reflect.TypeOf(Car{})

	for i := 0; i < tags.NumField(); i++ {
		field := tags.Field(i)
		fmt.Println(field.Tag, field.Name, field.Type)
	}

	car1 := Car{
		"520D",
		"Silver",
		"BMW",
		spec{4000, 1000, 2000},
	}
	fmt.Printf("%#v\n", car1)
	fmt.Printf("%#v\n", car1.detail)
}
