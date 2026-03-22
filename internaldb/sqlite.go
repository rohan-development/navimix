package internaldb

import (
	"database/sql"
	"encoding/json"

	_ "modernc.org/sqlite"
)

const db_name = "NaviMix.db"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite", db_name)
	check_err(err)
	//	defer db.Close()
	query := `
CREATE TABLE IF NOT EXISTS elements (
	id INTEGER PRIMARY KEY,
    data TEXT
);
`
	_, err = db.Exec(query)
	check_err(err)
}

func Add_element(id int, element any) {
	// db, err := sql.Open("sqlite", db_name)
	// check_err(err)
	// defer db.Close()
	jsonData, err := json.Marshal(element)
	check_err(err)
	_, err = db.Exec("INSERT INTO elements (id, data) VALUES (?, ?)", id, jsonData)
	check_err(err)
}

func Element_exists(id int) bool {
	// db, err := sql.Open("sqlite", db_name)
	// check_err(err)
	// defer db.Close()
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM elements WHERE id = ?)"
	err := db.QueryRow(query, id).Scan(&exists)
	check_err(err)
	return exists
}
