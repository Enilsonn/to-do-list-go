package repository

import (
	"context"
	"database/sql"
	"log"
)

func InitSQLiteDB(ctx context.Context, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	createTableTaskQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT NOT NULL,
		"done" BOOLEAN NOT NULL DEFAULT 0
	);`

	if _, err = db.ExecContext(ctx, createTableTaskQuery); err != nil {
		return nil, err
	}

	log.Println("database stable connection!")
	return db, nil
}
