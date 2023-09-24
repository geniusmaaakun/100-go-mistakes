package main

import (
	"context"
	"net/http"
	"time"

	"github.com/teivah/100-go-mistakes/08-concurrency-foundations/60-contexts/flight"
)

type publisher interface {
	Publish(ctx context.Context, position flight.Position) error
}

type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(position flight.Position) error {
	// タイムアウトを設定
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// キャンセル通知
	defer cancel()
	return h.pub.Publish(ctx, position)
}

// key は公開されていない独自の型を使うと良い。衝突を避けるため
type key string

const isValidHostKey key = "isValidHost"

func checkValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validHost := r.Host == "acme"
		// コンテキストに値を設定
		ctx := context.WithValue(r.Context(), isValidHostKey, validHost)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(ctx context.Context, ch chan Message) error {
	for {
		select {
		case msg := <-ch:
			// Do something with msg
			_ = msg
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

type Message struct{}
