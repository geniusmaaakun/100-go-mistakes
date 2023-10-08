package main

import "database/sql"

var dsn = ""

func listing1() error {
	//必ず確立するわけではない
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 確率を保証するためにPingを呼ぶ
	if err := db.Ping(); err != nil {
		return err
	}

	_ = db
	return nil
}
