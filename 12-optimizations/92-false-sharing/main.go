package main

import "sync"

type Input struct {
	a int64
	b int64
}

// input a, bの総和を求める
type Result1 struct {
	sumA int64
	sumB int64
}

//　偽共有が起きる
func count1(inputs []Input) Result1 {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result1{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}

// パディングを入れることで偽共有を防ぐ
// こっちの方が高速化される
type Result2 struct {
	sumA int64
	_    [56]byte
	sumB int64
}

func count2(inputs []Input) Result2 {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result2{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}
