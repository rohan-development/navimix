package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const db_name = "NaviMix.db"

func init_db() {
	db, err := sql.Open("sqlite3", db_name)
	check_err(err)
	defer db.Close()
	query := `
CREATE TABLE IF NOT EXISTS songs (
    id_deezer INTEGER PRIMARY KEY,
	id_navidrome TEXT,
    title TEXT,
    album TEXT,
	artist TEXT,
	album_id INTEGER,
	artist_id INTEGER,
	album_exists BOOLEAN
);
`
	_, err = db.Exec(query)
	check_err(err)
}

func add_song_to_database(song Song) {
	db, err := sql.Open("sqlite3", db_name)
	check_err(err)
	defer db.Close()
}

func song_exists_in_db(id int) bool {
	db, err := sql.Open("sqlite3", db_name)
	check_err(err)
	defer db.Close()
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM songs WHERE id = ?)"
	err = db.QueryRow(query, id).Scan(&exists)
	return exists
}
