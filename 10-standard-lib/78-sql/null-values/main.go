package main

import "database/sql"

//値がない場合はnilにならない、エラーになる
func listing1(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return err
	}
	// Defer closing rows

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}
	return nil
}

//値がない場合はnilになる
func listing2(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return err
	}
	// Defer closing rows

	var (
		//ポインタにする
		department *string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}
	return nil
}

//値がない場合はnilになる
func listing3(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return err
	}
	// Defer closing rows

	var (
		// NullStringはsql.NullString型
		department sql.NullString
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}
	return nil
}
