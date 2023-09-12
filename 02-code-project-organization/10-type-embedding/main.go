package main

import (
	"io"
	"os"
	"sync"
)

type Foo struct {
	Bar
}

type Bar struct {
	Baz int
}

func fooBar() {
	foo := Foo{}
	foo.Baz = 42
}

type InMem struct {
	// 間違った埋め込みmutexが公開されてしまう
	//sync.Mutex

	// 正しい埋め込み
	// 外部からアクセスできない。通常は内部でしか使わないので、
	mu sync.Mutex
	m  map[string]int
}

func New() *InMem {
	return &InMem{m: make(map[string]int)}
}

func (i *InMem) Get(key string) (int, bool) {
	//i.Lock()
	i.mu.Lock()
	v, contains := i.m[key]
	//i.Unlock()
	i.mu.Unlock()
	return v, contains
}

//正しい埋め込み
type Logger struct {
	// これは間違っている
	//writeCloser io.WriteCloser

	//埋め込むと、interfaceを満たす
	io.WriteCloser
}

// interfaceを満たす
// func (l Logger) Write(p []byte) (int, error) {
// 	return l.writeCloser.Write(p)
// }

// // interfaceを満たす
// func (l Logger) Close() error {
// 	return l.writeCloser.Close()
// }

func main() {
	//l := Logger{writeCloser: os.Stdout}
	l := Logger{WriteCloser: os.Stdout}

	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
