package main

import (
	"fmt"
)

//nil empty ¥sliceの使い分け
//　割当をしない必要としない場合はnilにする。空スライスを返す場合は空スライスを返す
//余分なメモリ割り当てをしないため。
// 初期要素を持つ場合は sl := []int{1, 2, 3} のように初期化する

/*
1: empty=true   nil=true
2: empty=true   nil=true
3: empty=true   nil=false
4: empty=true   nil=false
*/

func main() {
	// nil slice
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	// empty slice
	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}
