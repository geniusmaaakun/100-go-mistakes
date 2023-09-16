package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte
}

func main() {
	// 1000 個の要素を持つスライス
	foos := make([]Foo, 1_000)

	//
	printAlloc()

	// 各要素に1 MBのスライスを割り当てる
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	// 2つの要素だけを保持するスライスを作成する
	//two := keepFirstTwoElementsOnly(foos)

	//意図的に解放する場合
	//two := keepFirstTwoElementsOnlyCopy(foos)
	two := keepFirstTwoElementsOnlyMarkNil(foos)

	// GCを実行する
	runtime.GC()
	printAlloc()

	//アクセスできるくない？
	//fmt.Println(foos[100])

	// 参照を保持しているため、スライスの中身は解放されない
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	return foos[:2]
}

//元のスライスのアドレスとは異なるため、元のスライスのメモリは解放される
//容量を維持しなくても良い場合
func keepFirstTwoElementsOnlyCopy(foos []Foo) []Foo {
	res := make([]Foo, 2)
	copy(res, foos)
	return res
}

//意図的に解放する場合 nilにする
//容量を維持したい場合
func keepFirstTwoElementsOnlyMarkNil(foos []Foo) []Foo {
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
