package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestMySQLIntegration(t *testing.T) {
	setupMySQL()
	defer teardownMySQL()

	// ...
}

// 事後処理
func createConnection(t *testing.T, dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.FailNow()
	}
	t.Cleanup(
		func() {
			_ = db.Close()
		})
	return db
}

// 事前準備
func TestMain(m *testing.M) {
	setupMySQL()
	code := m.Run()
	teardownMySQL()
	os.Exit(code)
}

func setupMySQL() {}

func teardownMySQL() {}
