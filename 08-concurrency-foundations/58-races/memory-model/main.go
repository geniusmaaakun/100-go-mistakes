package main

import "fmt"

// i == 0
func listing1() {
	i := 0
	go func() {
		i++
	}()
}

// i == 0
func listing2() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}

// i == 1
func listing3() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	ch <- struct{}{}
}

// i == 1
func listing4() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	close(ch)
}

// i == 0 ブロックされない
func listing5() {
	i := 0
	ch := make(chan struct{}, 1)
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}

// i == 1
func listing6() {
	i := 0
	ch := make(chan struct{})
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}
