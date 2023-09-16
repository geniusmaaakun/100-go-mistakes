package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Init
	// メモリの確保
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	// Add elements
	// 要素の追加
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()

	// Remove elements
	//　要素の削除
	//　削除してもメモリは解放されない。マップのサイズは縮小されない
	// 一時的に容量が増えた場合、その容量は縮小されないので、定期的に再生成すると良い　欠点は瞬間的に二倍のメモリが必要になる。再配置するため
	// もう一つの解決法は、ポインタを保持することで、要素のメモリは小さくなる。64bitPCの場合、8バイトのポインタを保持するだけで、128バイトの要素を保持できる
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAlloc()

	/*
		KeepAlive は、その引数を現在到達可能としてマークします。 これにより、プログラム内で KeepAlive が呼び出される時点よりも前に、オブジェクトが解放されず、そのファイナライザーが実行されなくなります。
	*/
	runtime.KeepAlive(m)
}

func randBytes() [128]byte {
	return [128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
