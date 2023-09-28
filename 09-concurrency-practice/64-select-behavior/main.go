package main

import (
	"fmt"
	"time"
)

func main() {
	messageCh := make(chan int, 10)
	// 慣習的に、終了を通知するチャネルは、型を struct{} にする
	// これは、メモリを消費しないから 0 byte であることを意味する
	// interface{} にすると、メモリを消費する
	disconnectCh := make(chan struct{})

	go listing1(messageCh, disconnectCh)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}
	disconnectCh <- struct{}{}
	time.Sleep(10 * time.Millisecond)
}

// 途中で中断される可能性がある
func listing1(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func listing2(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)

		// ここで、disconnectCh が受信されるまで、messageCh からの受信をブロックする
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}
