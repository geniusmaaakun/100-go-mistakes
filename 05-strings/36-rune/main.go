package main

import "fmt"

func main() {
	// シングルバイト文字
	s := "hello"
	fmt.Println(len(s))

	//3byte文字　マルチバイト文字
	s = "汉"
	fmt.Println(len(s))

	s = string([]byte{0xE6, 0xB1, 0x89})
	fmt.Printf("%s\n", s)
}
