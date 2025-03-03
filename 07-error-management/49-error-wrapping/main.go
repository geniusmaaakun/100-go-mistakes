package main

import "fmt"

func bar() error {
	return barError{}
}

// ラップ機能がない時代はカスタムエラーを作成していた
type barError struct{}

func (b barError) Error() string {
	return "bar error"
}

func listing1() error {
	err := bar()
	if err != nil {
		return err
	}
	// ...
	return nil
}

///カスタムエラー
type BarError struct {
	Err error
}

func (b BarError) Error() string {
	return "bar failed:" + b.Err.Error()
}

func listing2() error {
	err := bar()
	if err != nil {
		return BarError{Err: err}
	}
	// ...
	return nil
}

//エラーをラップする
func listing3() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("bar failed: %w", err)
	}
	// ...
	return nil
}

//ラップされていない
func listing4() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("bar failed: %v", err)
	}
	// ...
	return nil
}
