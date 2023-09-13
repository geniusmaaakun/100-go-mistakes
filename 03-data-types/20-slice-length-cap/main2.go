package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:3]

	fmt.Println(cap(sl))
	//配列の4に位置が上書きされる
	//容量を超えた場合は、別の配列を参照することになる
	sl = append(sl, 6)

	fmt.Println(sl)
	fmt.Print(arr)

	/*
		4
		[2 3 6]
		[1 2 3 6 5
	*/
}
