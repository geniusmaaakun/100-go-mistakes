package main

import "fmt"

type customer struct {
	balance float64
}

// addはレシーバーの値を変更するため、ポインタレシーバーを使う
// 変更されない
func (c customer) add(v float64) {
	c.balance += v
}

func main() {
	c := customer{balance: 100.}
	c.add(50.)
	fmt.Printf("balance: %.2f\n", c.balance)
}
