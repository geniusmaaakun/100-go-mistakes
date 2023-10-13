//go:build integration
// +build integration

//タグを使ってテストを分類する
// 通常のテストを実行するには、タグを指定しないで go test を実行します。この場合、タグが付いていないテストだけが実行されます。
//go test -tags=integration ./...

package db

import (
	"os"
	"testing"
)

func TestInsert1(t *testing.T) {
	// ...
}

// 環境変数を使ってテストを分類する
// スキップされる
func TestInsert2(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test")
	}

	// ...
}
