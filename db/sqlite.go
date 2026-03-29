package db

import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)

const db_name = "navimix.db"

var db *sql.DB
var writeMutex sync.Mutex

func init() {
	var err error
	db, err = sql.Open("sqlite", db_name)
	check_err(err)
	//	defer db.Close()
	query := `
CREATE TABLE IF NOT EXISTS tracks (
	id INTEGER PRIMARY KEY,
	deezer_id INTEGER UNIQUE,
    title TEXT,
	album TEXT,
	artist TEXT,
	albumID TEXT,
	duration INTEGER
);
CREATE TABLE IF NOT EXISTS albums (
	id INTEGER PRIMARY KEY,
	deezer_id TEXT UNIQUE,
    title TEXT,
	artist TEXT,
	genre TEXT,
	year INTEGER
);
CREATE TABLE IF NOT EXISTS artists (
	id INTEGER PRIMARY KEY,
    data TEXT
);
`
	_, err = db.Exec(query)
	check_err(err)
}
