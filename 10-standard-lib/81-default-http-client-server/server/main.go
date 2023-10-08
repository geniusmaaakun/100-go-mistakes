package main

import (
	"net/http"
	"time"
)

// サーバー側でタイムアウトを設定する

func main() {
	s := &http.Server{
		Addr: ":8080",
		// 必須
		ReadHeaderTimeout: 500 * time.Millisecond, //ヘッダーの読み込みタイムアウト
		ReadTimeout:       500 * time.Millisecond, //ボディの読み込みタイムアウト
		// ハンドラをラップする
		Handler: http.TimeoutHandler(handler{}, time.Second, "foo"),

		// KeepAliveが有効な場合、次のリクエストまでの最大時間を設定できる
		IdleTimeout: 500 * time.Millisecond, //アイドルタイムアウト
	}
	_ = s
}

type handler struct{}

func (h handler) ServeHTTP(http.ResponseWriter, *http.Request) {}
