package db

import (
	"navimix/deezer"
)

func AddAlbum(album deezer.Album) {
	query := `
	INSERT INTO albums (deezer_id, title, artist, genre,
	year) VALUES (?,?,?,?,?) ON CONFLICT(deezer_id)
	DO NOTHING
	`
	if len(album.Genres.Data) > 0 {
		_, err := db.Exec(query, album.ID, album.Name,
			album.Artist.Name, album.Genres.Data[0].Name, album.Year)
		check_err(err)
	}
}

func GetAlbum(deezer_id string) (deezer.Album, error) {
	var a deezer.Album
	//var genre deezer.Genre
	a.Genres.Data = append(a.Genres.Data, deezer.Genre{})
	err := db.QueryRow(`
        SELECT deezer_id, title, artist, genre, year
        FROM albums WHERE deezer_id = ?
    `, deezer_id).Scan(
		&a.ID,
		&a.Name,
		&a.Artist.Name,
		&a.Genres.Data[0].Name,
		&a.Year,
	)
	//a.Genres.Data[0] = genre.Name
	if err == nil {
		return a, nil
	}
	return deezer.Album{}, err
}
