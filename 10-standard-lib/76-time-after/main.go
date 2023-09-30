package main

import (
	"context"
	"log"
	"time"
)

//メモリリークの原因になる
func consumer1(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
		case <-time.After(time.Hour):
			log.Println("warning: no messages received")
		}
	}
}

//メモリリークの原因にならない
// 欠点としては毎回contextを作成するので、パフォーマンスが悪い
func consumer2(ch <-chan Event) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			//メッセージを受信したらcontextをキャンセルする
			cancel()
			handle(event)
			//contextがキャンセルされた表示される
		case <-ctx.Done():
			log.Println("warning: no messages received")
		}
	}
}

//メモリリークの原因にならない
func consumer3(ch <-chan Event) {
	timerDuration := 1 * time.Hour
	timer := time.NewTimer(timerDuration)

	// 本番コードでは無限ループにすべきでない
	// cancelできるようにする
	for {
		//
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			handle(event)
		// タイムアウトしたらメッセージを表示する
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}

type Event struct{}

func handle(Event) {
}
