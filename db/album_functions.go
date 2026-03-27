package db

import "navimix/types"

func AddAlbum(album types.Album) {
	query := `
	INSERT INTO tracks (deezer_id, title, artist, genre,
	year) VALUES (?,?,?,?,?) ON CONFLICT(deezer_id) 
	DO NOTHING
	`
	_, err := db.Exec(query, album.ID, album.Title,
		album.Artist, album.Genre, album.Year)
	check_err(err)
}

func GetAlbum(deezer_id string) (types.Album, error) {
	var a types.Album
	err := db.QueryRow(`
        SELECT deezer_id, title, artist, genre, year
        FROM album WHERE deezer_id = ?
    `, deezer_id).Scan(
		&a.ID,
		&a.Title,
		&a.Artist,
		&a.Genre,
		&a.Year,
	)
	if err == nil {
		return a, nil
	}
	return types.Album{}, err
}
