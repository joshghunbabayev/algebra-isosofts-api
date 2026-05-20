package database

import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	instance *sql.DB
	once     sync.Once
)

func GetDatabase() *sql.DB {
	once.Do(func() {
		db, err := sql.Open("sqlite", "./database/main.db")
		if err != nil {
			panic(err)
		}
		db.Exec("PRAGMA journal_mode=WAL")
		db.Exec("PRAGMA busy_timeout=5000")
		instance = db
	})
	return instance
}
