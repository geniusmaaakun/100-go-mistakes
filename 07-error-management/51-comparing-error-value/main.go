package main

import (
	"database/sql"
	"errors"
)

func listing1() {
	err := query()
	if err != nil {
		// ラップされた場合falseになってしまう
		if err == sql.ErrNoRows {
			// ...
		} else {
			// ...
		}
	}
}

func listing2() {
	err := query()
	if err != nil {
		//　ラップされていてもtrueになる
		// 当れられた値 err と　再起的に値を比較する
		if errors.Is(err, sql.ErrNoRows) {
			// ...
		} else {
			// ...
		}
	}
}

func query() error {
	return nil
}
