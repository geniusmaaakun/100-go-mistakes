package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// 部分文字列を取り出す
	s1 := "Hello, World!"
	s2 := s1[:5]
	fmt.Println(s2)

	//複数バイト文字の場合はruneにする
	s1 = "Hêllo, World!"
	s2 = string([]rune(s1)[:5])
	fmt.Println(s2)
}

//メモリリークの例
type store struct{}

func (s store) handleLog1(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	// 36文字目までをuuidとして扱うが、logが巨大な場合は、メモリリークの原因になる。GCで回収されない
	uuid := log[:36]
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog2(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	//新たな文字列を基底配列から作成する
	uuid := string([]byte(log[:36]))
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog3(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	// 1.18から使えるようになったCloneを使う
	uuid := string(strings.Clone(log[:36]))
	s.store(uuid)
	// Do something
	return nil
}

func (s store) store(uuid string) {
	// ...
}
