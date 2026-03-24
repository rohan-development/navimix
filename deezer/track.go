package deezer

import (
	"encoding/json"
	"io"
	"navimix/api"
	"navimix/deezer"
	"net/http"
	"strconv"
	"strings"
)

func GetTrack(query string) Data {
	url := deezer_api_base + "artist/" + query
	response, err := http.Get(url)
	check_err(err)
	data, err := io.ReadAll(response.Body)
	var link Data
	err = json.Unmarshal(data, &link)
	check_err(err)
	return link
}

func ExtractTrack(main_search api.SubsonicResponse,
	deezer_search []deezer.Data, num_songs, search_version int) SubsonicResponse {
	var add_song Song
	var add_album Song
	var add_artist Artist
	//var add_artist Artist
	delta := 0
	for i := 0; i+delta < num_songs; i += 1 {
		switch deezer_search[i+delta].Type {
		case "track":
			add_song.ID = strconv.Itoa(deezer_search[i+delta].ID)
			add_song.Title = deezer_search[i+delta].Title
			add_song.Album = deezer_search[i+delta].Album.Name
			add_song.Artist = deezer_search[i+delta].Artist.Name
			add_song.AlbumID = strconv.Itoa(deezer_search[i+delta].Album.ID)
			add_song.Parent = add_song.AlbumID
			add_song.BitRate = 128
			add_song.ContentType = "audio/mp3"
			add_song.Suffix = "mp3"
			// album := query_deezer_api("album/" + add_song.AlbumID)
			// add_song.Year, _ = strconv.Atoi(album.Year[0:4])
			// add_song.Genre = album.Genres.Data[0].Name
			add_song.Duration = deezer_search[i+delta].Duration
			add_song.SortName = strings.ToLower(add_song.Title)
			add_song.Type = "music"
			add_song.MediaType = "song"
			add_song.DisplayArtist = add_song.Artist
			if song_in_search(main_search, add_song, search_version) {
				delta += 1
				i -= 1
				continue
			}

			if search_version == 2 {
				main_search.SearchResult2.Song = append(main_search.SearchResult2.Song,
					add_song)
			} else {
				main_search.SearchResult3.Song = append(main_search.SearchResult3.Song,
					add_song)
			}
		case "album":
			add_album.ID = strconv.Itoa(deezer_search[i+delta].ID)
			add_album.Title = deezer_search[i+delta].Title
			add_album.Artist = deezer_search[i+delta].Artist.Name
			add_album.Parent = strconv.Itoa(deezer_search[i+delta].Artist.ID)
			add_album.MediaType = "album"
			add_album.DisplayArtist = add_album.Artist
			// album := query_deezer_api("album/" + add_album.ID)
			// add_album.Year, _ = strconv.Atoi(album.Year[0:4])
			// add_album.Genre = album.Genres.Data[0].Name
			if album_in_search(main_search, add_album, search_version) {
				delta += 1
				i -= 1
				continue
			}
			if search_version == 2 {
				main_search.SearchResult2.Album = append(main_search.SearchResult2.Album,
					add_album)
			} else {
				main_search.SearchResult3.Album = append(main_search.SearchResult3.Album,
					add_album)
			}
		case "artist":
			add_artist.ID = strconv.Itoa(deezer_search[i+delta].ID)
			add_artist.Name = deezer_search[i+delta].Name
			add_artist.CoverArt = deezer_search[i+delta].ArtistBig
			if artist_in_search(main_search, add_artist, search_version) {
				delta += 1
				i -= 1
				continue
			}
			if search_version == 2 {
				main_search.SearchResult2.Artist = append(main_search.SearchResult2.Artist,
					add_artist)
			} else {
				main_search.SearchResult3.Artist = append(main_search.SearchResult3.Artist,
					add_artist)
			}
		}

	}

	//add_song.Artists[0].Name = add_song.Artist
	// if !internaldb.Element_exists(deezer_search[i+delta].ID) { //add track to db
	// 	internaldb.Add_element(deezer_search[i+delta].ID, add_song)
	// }
	// if !internaldb.Element_exists(deezer_search[i+delta].Album.ID) { //add album to db
	// 	internaldb.Add_element(deezer_search[i+delta].ID, deezer_search[i+delta].Album)
	// }

	// if !song_exists_in_db(deezer_search[i+delta].ID) {
	// 	add_song_to_database(add_song)
	// }
	return main_search

}
