package main

import "strings"

//遅い
func concat1(values []string) string {
	s := ""
	for _, value := range values {
		//反復ごとに新しい文字列を割り当てる。非効率
		s += value
	}
	return s
}

//strings.Builderを使う
func concat2(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		//エラーを返すことはない
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

//先に容量確保するので高速
func concat3(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}
