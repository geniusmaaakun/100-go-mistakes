package main

// 構造体のスライス
type Foo struct {
	a int64
	b int64
}

func sumFoo(foos []Foo) int64 {
	var total int64
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}
	return total
}

//スライスの構造体
// スライスの構造体はコンパクトなので、反復処理に必要なキャッシュラインの数が少ない
// こっちの方が20%早い
type Bar struct {
	a []int64
	b []int64
}

func sumBar(bar Bar) int64 {
	var total int64
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}
