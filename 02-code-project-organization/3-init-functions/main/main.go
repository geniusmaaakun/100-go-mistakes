package main

import (
	"fmt"

	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/redis"
)

//initの実行順序は以下の通り

//また複数のinitがある場合、ファイル名順に実行される

//1
func init() {
	fmt.Println("init 1")
}

//2
func init() {
	fmt.Println("init 2")
}

//3
func main() {
	err := redis.Store("foo", "bar")
	_ = err
}
