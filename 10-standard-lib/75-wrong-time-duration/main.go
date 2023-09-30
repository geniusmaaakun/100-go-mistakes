package main

import (
	"fmt"
	"time"
)

func main() {
	listing1()
	//listing2()
}

func listing1() {
	// 1ms = 1000us
	ticker := time.NewTicker(1000)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}

func listing2() {
	// 1microsecond = の場合はこのようにした方が良い
	ticker := time.NewTicker(time.Microsecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}
