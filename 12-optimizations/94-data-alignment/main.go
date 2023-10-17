package main

type Foo1 struct {
	b1 byte  // 1バイト
	i  int64 // 8バイト
	b2 byte  // 1バイト
}

func sum1(foos []Foo1) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}

//こっちの方が高速化される
type Foo2 struct {
	i  int64 // 8バイト
	b1 byte  // 1バイト
	b2 byte  // 1バイト
}

func sum2(foos []Foo2) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}
