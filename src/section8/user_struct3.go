package main

import "fmt"

type Cart struct {
	amount, price int
}

func main() {

	cart := Cart{amount: 10, price: 1000}

	fmt.Println(cart.Purchase())
	cart.TrySetAmount(20)
	fmt.Println("after TrySetAmount", cart.amount)
	cart.SetAmount(20)
	fmt.Println("after SetAmount", cart.amount)

}

func (cart Cart) Purchase() int {
	return cart.amount * cart.price
}

// 호출하는 쪽에서 &cart 로 넘기지 않아도 알아서 처리됨
func (cart *Cart) SetAmount(newAmount int) {
	cart.amount = newAmount
}

func (cart Cart) TrySetAmount(newAmount int) {
	cart.amount = newAmount
}
