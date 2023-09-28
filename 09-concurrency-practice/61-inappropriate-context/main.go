package main

import (
	"context"
	"net/http"
	"time"
)

//問題あり
// 巻き戻しを行うっても、publish() が実行されてしまう
func handler1(w http.ResponseWriter, r *http.Request) {
	// Http レスポンスを計算するためのタスクを実行
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 非同期で処理
	// contextがキャンセルされたら、メッセージ発行を中断する
	go func() {
		err := publish(r.Context(), response)
		// Do something with err
		_ = err
	}()

	// Http レスポンスを書き込む
	writeResponse(response)
}

//
func handler2(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		// 空のコンテキストを渡す
		// 親コンテキストを伝播させない
		// ただし親コンテキストで有用な値が含まれている場合は、それらの値を失うことになる
		err := publish(context.Background(), response)
		// Do something with err
		_ = err
	}()

	writeResponse(response)
}

//改善
func handler3(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {

		// キャンセルシグナルをを伝えないように、独自のコンテキストを作成
		err := publish(detach{ctx: r.Context()}, response)
		// Do something with err
		_ = err
	}()

	writeResponse(response)
}

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}

func doSomeTask(context.Context, *http.Request) (string, error) {
	return "", nil
}

func publish(context.Context, string) error {
	return nil
}

func writeResponse(string) {}
