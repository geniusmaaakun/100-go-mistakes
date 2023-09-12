package main

import (
	"fmt"
	"sort"
)

// 非推奨
func getKeys(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		// ...
	}

	return nil, nil
}

// 推奨
// genericsを使って改善
// keyはanyでなく、comparableである必要がある
func getKeysGenerics[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// カスタム制約
type customConstraint interface {
	~int | ~string
}

// カスタム制約を使って改善
func getKeysWithConstraint[K customConstraint, V any](m map[K]V) []K {
	// ...
	return nil
}

//型パラメーターの制約
type Node[T any] struct {
	Val  T
	next *Node[T]
}

// レシーバーにも型情報を渡す
// メソッドには型パラメーターを指定できない
func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

//一般的な用途
type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

//dataInterfaceを実装していることを確認する
func (s SliceFn[T]) Len() int           { return len(s.S) }
func (s SliceFn[T]) Less(i, j int) bool { return s.Compare(s.S[i], s.S[j]) }
func (s SliceFn[T]) Swap(i, j int)      { s.S[i], s.S[j] = s.S[j], s.S[i] }

func main() {
	s := SliceFn[int]{
		S: []int{3, 2, 1},
		Compare: func(a, b int) bool {
			return a < b
		},
	}
	sort.Sort(s)
	fmt.Println(s.S)
}
