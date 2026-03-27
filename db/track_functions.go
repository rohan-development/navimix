package db

import "navimix/types"

func AddTrack(track types.Song) {
	query := `
	INSERT INTO tracks (deezer_id, title, album, artist,
	albumID, duration) VALUES (?,?,?,?,?,?) ON CONFLICT(deezer_id) 
	DO NOTHING
	`
	_, err := db.Exec(query, track.ID, track.Title, track.Album,
		track.Artist, track.AlbumID, track.Duration)
	check_err(err)
}

func GetTrack(deezer_id string) (types.Song, error) {
	var t types.Song
	err := db.QueryRow(`
        SELECT deezer_id, title, album, artist, albumID, duration
        FROM tracks WHERE deezer_id = ?
    `, deezer_id).Scan(
		&t.ID,
		&t.Title,
		&t.Album,
		&t.Artist,
		&t.AlbumID,
		&t.Duration,
	)
	if err == nil {
		return t, nil
	}
	return types.Song{}, err
}
