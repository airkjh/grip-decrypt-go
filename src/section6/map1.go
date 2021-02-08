package main

import "fmt"

func main() {
	// Reference Type
	// 참조 타입이므로 비교 연산자 사용 불가
	// 참조타입은 key 로 사용 불가, value 로는 모든 타입 사용 가능
	// make 함수 및 리터럴 초기화 가능
	// 삽입한 순서가 보장되지 않음

	var map1 map[string]int = make(map[string]int)
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	map4 := map[string]int{
		"A": 1,
		"B": 2,
	}

	fmt.Println(map4)

	map5 := map[string]int{}
	map5["apple"] = 15
	map5["banana"] = 400
	map5["orange"] = 33

	fmt.Println(map5)

	map6 := make(map[string]int, 10)
	map6["apple"] = 15
	map6["banana"] = 400
	map6["orange"] = 33
	fmt.Println(map6)
	fmt.Println(map6["orange"])

	for key, value := range map5 {
		fmt.Println("key=", key, "value=", value)
	}

	map7 := map[string]string{
		"google": "Google",
		"daum":   "Daum",
		"naver":  "Naver",
	}

	map7["google"] = "Gooogle"

	fmt.Println(map7)

	delete(map7, "daum")

	fmt.Println(map7)

	// 조회할 경우 주의할 점
	//
	map8 := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 30,
		"lemon":  0,
	}

	value1 := map8["lemon"]
	value2 := map8["kiwi"]           // 존재하지 않는 키
	value3, keyExist := map8["wiki"] // keyExist <- 키가 맵에 존재하는지?

	fmt.Println(value1)
	fmt.Println(value2)           // 키가 존재하지 않을 때는 value 타입의 기본 값
	fmt.Println(value3, keyExist) // 키가 존재하지 않아서 0인지 실제값이 0인지 구분하기 위해 두 번째 값 (키 존재 여부)를 함께 리턴해준다

	if _, exists := map8["kiwi"]; !exists {
		fmt.Println("Kiwi does not exists")
	}
}
