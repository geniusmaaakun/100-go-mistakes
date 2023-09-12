package main

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

// dbの初期化をinitでやるべきではない
// initはエラーを返さず、パニックを起こすことしかできない
// テストがやりにくい
// グローバル変数を使うのはよくない　→　変更ができてしまう
//初期化関数を別で作成し、カプセル化する  (*sql.DB, error)を返す
// 本ではinitを使っているが、initは使わない方がいいことの方が多い
// P18を参照
func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = d
}

func createClient(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
