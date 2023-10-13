package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// サーバーのテスト
func TestHandler(t *testing.T) {
	// リクエストを作成
	req := httptest.NewRequest(http.MethodGet, "http://localhost",
		strings.NewReader("foo"))
	//レスピンスレコーダーを作成
	w := httptest.NewRecorder()
	// ハンドラ実行
	Handler(w, req)

	//検証
	if got := w.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	body, _ := ioutil.ReadAll(w.Body)
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	if http.StatusOK != w.Result().StatusCode {
		t.FailNow()
	}
}

// クライアントのテスト
func TestDurationClientGet(t *testing.T) {
	//サーバーを作成
	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"duration": 314}`))
			},
		),
	)
	defer srv.Close()

	client := NewDurationClient()
	duration, err :=
		client.GetDuration(srv.URL, 51.551261, -0.1221146, 51.57, -0.13)
	if err != nil {
		t.Fatal(err)
	}

	if duration != 314*time.Second {
		t.Errorf("expected 314 seconds, got %v", duration)
	}
}
