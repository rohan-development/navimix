package db

import (
	"navimix/deezer"
	"strconv"
)

func AddTrack(track deezer.Data) {
	writeMutex.Lock()
	defer writeMutex.Unlock()
	query := `
	INSERT INTO tracks (deezer_id, title, album, artist,
	albumID, duration) VALUES (?,?,?,?,?,?) ON CONFLICT(deezer_id) 
	DO NOTHING
	`
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA busy_timeout=5000;") // wait up to 5 seconds if locked
	_, err := db.Exec(query, track.ID, track.Title, track.Album.Name,
		track.Artist.Name, strconv.Itoa(track.Album.ID), track.Duration)
	check_err(err)
}

func GetTrack(deezer_id string) (deezer.Data, error) {
	var t deezer.Data
	var albumID string
	err := db.QueryRow(`
        SELECT deezer_id, title, album, artist, albumID, duration
        FROM tracks WHERE deezer_id = ?
    `, deezer_id).Scan(
		&t.ID,
		&t.Title,
		&t.Album.Name,
		&t.Artist.Name,
		&albumID,
		&t.Duration,
	)
	t.Album.ID, _ = strconv.Atoi(albumID)
	if err == nil {
		return t, nil
	}
	return deezer.Data{}, err
}
