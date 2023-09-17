package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s {
		// s[i]ではbyteになる
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))

	//解決策
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	//解決策
	//runeに変換。オーバーヘッドがある
	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}

	s2 := "hello"
	fmt.Printf("%c\n", rune(s2[4]))
}

func getIthRune(largeString string, i int) rune {
	for idx, v := range largeString {
		if idx == i {
			return v
		}
	}
	return -1
}
