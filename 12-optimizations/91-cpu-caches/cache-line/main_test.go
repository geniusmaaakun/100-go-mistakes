package main

import "testing"

// 4倍の速度差があると思うが、実際は10%ほどしか差がない

var global int64

func BenchmarkSum2(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum2(s)
	}
	global = local
}

func BenchmarkSum8(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum8(s)
	}
	global = local
}
