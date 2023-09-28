package main

import (
	"fmt"
	"sync"
)

func main() {
	c := Cache{
		balances: make(map[string]float64),
	}
	c.AddBalance("1", 1.0)
	c.AddBalance("2", 3.0)
	fmt.Println(c.AverageBalance1())
	fmt.Println(c.AverageBalance2())
	fmt.Println(c.AverageBalance3())
}

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

//別のゴルーチンで更新をしている場合、データ競合を引き起こす
func (c *Cache) AverageBalance1() float64 {
	// blancesを代入するときだけ、ロックする
	c.mu.RLock()
	//　参照なので、コピーとは異なる
	balances := c.balances
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	return sum / float64(len(balances))
}

//競合しない
func (c *Cache) AverageBalance2() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(c.balances))
}

// 競合しない
func (c *Cache) AverageBalance3() float64 {
	c.mu.RLock()
	// ロックを解除する前に、コピーを作成する
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}
