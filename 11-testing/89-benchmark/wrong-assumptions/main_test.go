package main

import (
	"sync/atomic"
	"testing"
)

//マイクロベンチマークについて間違った仮定をする

//誤差のような差がある場合は、ベンチマークの結果は信頼できない

func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}
