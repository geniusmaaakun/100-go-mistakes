package main

import "fmt"

//setter, getter例
//特に必須ではないが、やるとしたらこんな感じ
// やるとするなら P20のケースを参照

type CustomerService struct {
	store string
}

// getStoreにしない
func (cs CustomerService) Store() string {
	return cs.store
}

func (cs *CustomerService) setStore(store string) {
	cs.store = store
}

func main() {
	cs := CustomerService{}
	cs.setStore("store")
	fmt.Println(cs.Store())
}
