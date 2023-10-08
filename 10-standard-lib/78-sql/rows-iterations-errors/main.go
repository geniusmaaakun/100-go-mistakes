package main

import (
	"context"
	"database/sql"
	"log"
)

//for rows.Next()でエラーが発生した場合はrows.Err()を呼ぶ必要がある
func get1(ctx context.Context, db *sql.DB, id string) (string, int, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}

	return department, age, nil
}

//改善版
func get2(ctx context.Context, db *sql.DB, id string) (string, int, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}

	// rows.Err()はrows.Next()の後に呼ぶ
	if err := rows.Err(); err != nil {
		return "", 0, err
	}

	return department, age, nil
}
