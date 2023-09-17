package main

import (
	"fmt"
	"strings"
)

func main() {
	// 含まれる限り取り除く
	fmt.Println(strings.TrimRight("123oxo", "xo"))

	// 末尾から文字セットを引くだけ
	fmt.Println(strings.TrimSuffix("123oxo", "xo"))

	// 含まれる限り取り除く
	fmt.Println(strings.TrimLeft("oxo123", "ox"))

	// 先頭から文字セットを引くだけ
	fmt.Println(strings.TrimPrefix("oxo123", "ox"))

	//left + right
	fmt.Println(strings.Trim("oxo123oxo", "ox"))
}
