package main

import (
	"database/sql"
	"log"
)

const query = "..."

// defer のエラーを無視してしまう例
func getBalance1(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// Use rows
	return 0, nil
}

// deferのエラーを無視しない例
func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }()

	// Use rows
	return 0, nil
}

// メッセージを出力する例
func getBalance3(db *sql.DB, clientID string) (balance float32, err error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil {
			if closeErr != nil {
				log.Printf("failed to close rows: %v", closeErr)
			}
			return
		}
		err = closeErr
	}()

	// Use rows
	return 0, nil
}
