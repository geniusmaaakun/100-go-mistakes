package main

import "fmt"

//予測不可能
func listing1() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	//　ソートされていないのでランダム
	// 毎回順序が異なる
	// 要素が追加された場合、それが呼び出されるか呼び出されないかは実行ごとに異なる
	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}

	fmt.Println(m)
}

//予測可能にする
func listing2() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	m2 := copyMap(m)

	// コピーしたマップを更新
	for k, v := range m {
		m2[k] = v
		if v {
			m2[10+k] = true
		}
	}

	fmt.Println(m2)
}

func copyMap(m map[int]bool) map[int]bool {
	res := make(map[int]bool, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

func main() {
	listing1()
	listing2()
}
