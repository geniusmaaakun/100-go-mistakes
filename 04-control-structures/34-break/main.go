package main

import (
	"context"
	"fmt"
)

// breakされない
func listing1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break
		}
	}
}

// breakされる
func listing2() {
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			// ラベル付きcontinue文で継続にすることもできる
			break loop
		}
	}
}

func listing3(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break
		}
	}
}

func listing4(ctx context.Context, ch <-chan int) {
loop:
	for {
		select {
		case <-ch:
			// Do something
		case <-ctx.Done():
			break loop
		}
	}
}
