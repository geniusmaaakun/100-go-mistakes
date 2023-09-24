package main

import "fmt"

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

func main() {
	_ = f1()
	_ = f2()
	_ = f3()
}

func f1() error {
	var status string

	//初期値が空文字列のため、deferで呼び出される時点では空文字列が渡される
	defer notify(status)
	defer incrementCounter(status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

//解決策1
func f2() error {
	var status string
	//deferで呼び出される時点でのstatusのポインタを渡す
	defer notifyPtr(&status)
	defer incrementCounterPtr(&status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

//解決策2
func f3() error {
	var status string
	//クロージャーにすることで最後に評価されたタイミングでのstatusを渡す
	defer func() {
		notify(status)
		incrementCounter(status)
	}()

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func notify(status string) {
	fmt.Println("notify:", status)
}

func incrementCounter(status string) {
	fmt.Println("increment:", status)
}

func notifyPtr(status *string) {
	fmt.Println("notify:", *status)
}

func incrementCounterPtr(status *string) {
	fmt.Println("increment:", *status)
}

func foo() error {
	return nil
}

func bar() error {
	return nil
}
