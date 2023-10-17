package main

import (
	"io"
	"sync"
)

// sync.Poolを使ってメモリの再利用をする
var pool = sync.Pool{
	New: func() any {
		return make([]byte, 1024)
	},
}

func write(w io.Writer) {
	// バッファを再利用する
	buffer := pool.Get().([]byte)
	// バッファを空にする。すでに使われている可能性があるため
	buffer = buffer[:0]
	// 使い終わったら返す
	defer pool.Put(buffer)

	// 何かを書き込む
	getResponse(buffer)
	// 書き込む
	_, _ = w.Write(buffer)
}

func getResponse([]byte) {
}
