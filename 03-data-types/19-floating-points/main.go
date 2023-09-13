package main

import "fmt"

//浮動小数点数の挙動
//近似値であることに注意　比較など

func main() {
	var n float32 = 1.0001
	fmt.Println(n * n)

	var a float64
	positiveInf := 1 / a
	negativeInf := -1 / a
	nan := a / a
	fmt.Println(positiveInf, negativeInf, nan)
}

//浮動小数点数の比較
//計算フローが異なるだけで答えが異なる
//nが大きくなるほどズレが大きくなる
//精度を上げるにはかけ算と割り算を先にやる
func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}
