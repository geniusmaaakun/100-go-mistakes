package main

type cache struct {
	m map[string]int
}

func (c *cache) get1(bytes []byte) (v int, contains bool) {
	key := string(bytes)
	v, contains = c.m[key]
	return
}

// コンパイラはbyteからstringへの変換を避けるために、高速化する
func (c *cache) get2(bytes []byte) (v int, contains bool) {
	v, contains = c.m[string(bytes)]
	return
}
