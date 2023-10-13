package main

import (
	"strings"
	"testing"
	"testing/iotest"
)

func TestLowerCaseReader(t *testing.T) {
	err := iotest.TestReader(
		// テスト対象のio.Reader
		&LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")},
		// 一致するかどうかを確認するための期待値
		[]byte("acegi"),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFoo1(t *testing.T) {
	// 2回目のReadAllでエラーが発生する
	err := foo1(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

func TestFoo2(t *testing.T) {
	// foo2で3回までリトライする
	err := foo2(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

func randomString(i int) string {
	return string(make([]byte, i))
}
