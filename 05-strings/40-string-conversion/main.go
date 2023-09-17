package main

import (
	"bytes"
	"io"
	"strings"
)

// 無駄な変換が入っている
func getBytes1(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// byteをstringに変換してからbyteに変換している
	return []byte(sanitize1(string(b))), nil
}

func sanitize1(s string) string {
	return strings.TrimSpace(s)
}

// 無駄な変換が入っていない
func getBytes2(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitize2(b), nil
}

func sanitize2(b []byte) []byte {
	return bytes.TrimSpace(b)
}
