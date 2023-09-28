package main

//二つのチャネルをマージする

//両方のチャネルから値を受信できない可能性がある
func merge1(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		// ch1からの受信が完了するまで、ブロックする
		for v := range ch1 {
			ch <- v
		}
		// ch2からの受信が完了するまで、ブロックする
		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

//両方のチャネルから値を受信できるがcloseされない
//ch1またはch2がcloseされたら、受信が永遠にされる。closeされたチャネルからの受信は、常にゼロ値を返す
func merge2(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

// メッセージとクローズを区別する
//このコードの問題点はどちらかがcloseされたら、ループが回り続けること
func merge3(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			// ch1がcloseされたら、ch1Closedをtrueにする
			case v, open := <-ch1:
				// closeされても、処理される
				if !open {
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()

	return ch
}

//nilチャネルを使う
func merge4(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					//  nilチャネルにすることで、select文からch1を除外する
					ch1 = nil
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}
