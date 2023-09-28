package main

import "fmt"

//競合しない
func listing1() {
	// 容量がいっぱいなので、新しいスライスを割り当てる
	s := make([]int, 1)

	go func() {
		// 新しいスライスを割り当てる
		s1 := append(s, 1)
		fmt.Println(s1)
	}()

	go func() {
		// 新しいスライスを割り当てる
		s2 := append(s, 1)
		fmt.Println(s2)
	}()
}

//コピーしてから追加するver
func listing2() {
	s := make([]int, 0, 1)

	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}()

	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s2 := append(sCopy, 1)
		fmt.Println(s2)
	}()
}
