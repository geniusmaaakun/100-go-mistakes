package main

import (
	"net"
	"net/http"
	"time"
)

//タイムアウトなどを設定する
// 終了しないリクエストでリソースが枯渇するのを防ぐ

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second, //全体的なタイムアウト
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second, //接続のタイムアウト
			}).DialContext,
			TLSHandshakeTimeout:   time.Second, //TLSのハンドシェイクのタイムアウト
			ResponseHeaderTimeout: time.Second, //レスポンスヘッダーのタイムアウト
		},
	}
	_ = client
}
