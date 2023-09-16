package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func main() {
	// 　初期化
	s := Store{
		m: make(map[string]*Customer),
	}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	print(s.m)
}

func (s *Store) storeCustomers(customers []Customer) {
	// customerのポインタを使い回すので　最後の要素のポインタを毎回更新していることになってしまう
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
}

func (s *Store) storeCustomers2(customers []Customer) {
	for _, customer := range customers {
		//新しく変数を定義しているので毎回新しいアドレスになる
		current := customer
		// アドレスをコピー
		s.m[current.ID] = &current
	}
}

func (s *Store) storeCustomers3(customers []Customer) {
	for i := range customers {
		customer := &customers[i]
		s.m[customer.ID] = customer
	}
}

func print(m map[string]*Customer) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%#v\n", k, v)
	}
}
