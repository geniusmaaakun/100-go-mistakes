package redis

import "fmt"

// init関数はmain関数よりも先に実行されます
//インポートされたときに実行される
func init() {
	fmt.Println("redis")
}

func Store(key, value string) error {
	return nil
}
