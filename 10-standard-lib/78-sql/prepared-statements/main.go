package main

import "database/sql"

func listing1(db *sql.DB, id string) error {
	// プリペアドステートメントを作成
	stmt, err := db.Prepare("SELECT * FROM ORDER WHERE ID = ?")
	if err != nil {
		return err
	}
	// クエリを実行
	rows, err := stmt.Query(id)
	if err != nil {
		return err
	}
	_ = rows

	//不要になったらクローズする
	//stmt.Close()
	return nil
}
