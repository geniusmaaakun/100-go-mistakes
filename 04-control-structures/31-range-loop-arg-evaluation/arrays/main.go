package main

import "fmt"

func listing1() {
	a := [3]int{0, 1, 2}
	// 評価時は0, 1, 2
	for i, v := range a {
		a[2] = 10
		// 2が出力される
		if i == 2 {
			// rangeで評価した値が表示される
			fmt.Println(v)
		}
	}
}

func listing2() {
	a := [3]int{0, 1, 2}
	for i := range a {
		a[2] = 10
		if i == 2 {
			// 10が出力される
			fmt.Println(a[2])
		}
	}
}

// ポインタなので高速
func listing3() {
	a := [3]int{0, 1, 2}
	//　ポインタなので更新される
	for i, v := range &a {
		a[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
}

func main() {
	listing1()
	listing2()
	listing3()
}
