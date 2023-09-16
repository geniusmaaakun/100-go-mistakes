package main

import "testing"

const n = 1_000_000

var global map[int]struct{}

//非効率なmapの初期化
// マップも倍々で容量を増やしていくので、最初に容量を指定しておくと効率的
//再配置を行うので、O(N）の計算量がかかる
func BenchmarkMapWithoutSize(b *testing.B) {
	var local map[int]struct{}
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{})
		for j := 0; j < n; j++ {
			m[j] = struct{}{}
		}
	}
	global = local
}

//効率的なmapの初期化
func BenchmarkMapWithSize(b *testing.B) {
	var local map[int]struct{}
	for i := 0; i < b.N; i++ {
		//初期サイズを指定することで、再配置を行わずに済む
		m := make(map[int]struct{}, n)
		for j := 0; j < n; j++ {
			m[j] = struct{}{}
		}
	}
	global = local
}
