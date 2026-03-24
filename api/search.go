package api

import (
	"encoding/json"
	"io"

	//"navimix/internaldb"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Search(writer http.ResponseWriter, req *http.Request) {
	search_version := 3
	if req.URL.Path[6:] == "search2.view" {
		search_version = 2
	}
	request := req.URL.Query().Get("query")
	// if len(request) > 0 {
	// 	request = request[1 : len(request)-1]
	// }
	main_search := Get_subsonic_response(writer, req, true)
	track_search := search_deezer(request, "") //empty string = track
	//var num_songs int
	num_songs, _ := strconv.Atoi(req.URL.Query().Get("songCount"))
	num_albums, _ := strconv.Atoi(req.URL.Query().Get("albumCount"))
	num_artists, _ := strconv.Atoi(req.URL.Query().Get("artistCount"))
	var album_search []deezer_data
	var artist_search []deezer_data

	if num_albums > 0 {
		album_search = search_deezer(request, "album")
	}

	if num_artists > 0 {
		artist_search = search_deezer(request, "artist")
	}

	if len(artist_search) < num_artists {
		num_artists = len(artist_search)
	}

	if len(album_search) < num_albums {
		num_albums = len(album_search)
	}
	if len(track_search) < num_songs {
		num_songs = len(track_search)
	}

	if main_search.StatusCode != "failed" {
		main_search = extract_deezer_elements(main_search, track_search, num_songs,
			search_version)
		// if search_version == 2 && num_albums > 0 {
		// 	main_search.SearchResult2.Album = extract_deezer_elements(main_search,
		// 		album_search, num_albums, search_version).SearchResult2.Album
		// } else if num_albums > 0 {
		// 	main_search.SearchResult3.Album = extract_deezer_elements(main_search,
		// 		album_search, num_albums, search_version).SearchResult3.Album
		// }

		if search_version == 2 && num_artists > 0 {
			main_search.SearchResult2.Artist = extract_deezer_elements(main_search,
				artist_search, num_artists, search_version).SearchResult2.Artist
		} else if num_artists > 0 {
			main_search.SearchResult3.Artist = extract_deezer_elements(main_search,
				artist_search, num_artists, search_version).SearchResult3.Artist
		}
	}
	var embedded EmbeddedResponse
	embedded.Subsonic = main_search
	//embedded.Subsonic.StatusCode = req.URL.Path[6:]
	json.NewEncoder(writer).Encode(embedded)
}

func search_deezer(query, attribute string) []deezer_data {
	url := deezer_search_base + attribute +
		"?q=" + url.QueryEscape(query)
	// if attribute == "all" {
	// 	url = deezer_search_base + url.QueryEscape(":"+query)
	// }
	response, err := http.Get(url)
	check_err(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	var link deezer_response
	err = json.Unmarshal(data, &link)
	check_err(err)
	return link.Data
}

func song_in_search(main_search SubsonicResponse, add_song Song,
	search_version int) bool {
	if search_version == 2 {
		for j := 0; j < len(main_search.SearchResult2.Song); j += 1 {
			if add_song.Title == main_search.SearchResult2.Song[j].Title &&
				add_song.Artist == main_search.SearchResult2.Song[j].Artist {
				return true
			}
		}
		return false
	}
	for j := 0; j < len(main_search.SearchResult3.Song); j += 1 {
		if add_song.Title == main_search.SearchResult3.Song[j].Title &&
			add_song.Artist == main_search.SearchResult3.Song[j].Artist {
			return true
		}
	}
	return false

}

func album_in_search(main_search SubsonicResponse, add_album Song,
	search_version int) bool {
	if search_version == 2 {
		for j := 0; j < len(main_search.SearchResult2.Album); j += 1 {
			if add_album.Title == main_search.SearchResult2.Album[j].Title &&
				add_album.Artist == main_search.SearchResult2.Album[j].Artist {
				return true
			}
		}
		return false
	}
	for j := 0; j < len(main_search.SearchResult3.Album); j += 1 {
		if add_album.Title == main_search.SearchResult3.Album[j].Title &&
			add_album.Artist == main_search.SearchResult3.Album[j].Artist {
			return true
		}
	}
	return false

}

func artist_in_search(main_search SubsonicResponse, add_artist Artist,
	search_version int) bool {
	if search_version == 2 {
		for j := 0; j < len(main_search.SearchResult2.Artist); j += 1 {
			if add_artist.Name == main_search.SearchResult2.Artist[j].Name {
				return true
			}
		}
		return false
	}
	for j := 0; j < len(main_search.SearchResult3.Artist); j += 1 {
		if add_artist.Name == main_search.SearchResult3.Artist[j].Name {
			return true
		}
	}
	return false

}

func extract_deezer_elements(main_search SubsonicResponse,
	deezer_search []deezer_data, num_songs, search_version int) SubsonicResponse {
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

// func Search3(writer http.ResponseWriter, req *http.Request) {
// 	request := req.URL.Query().Get("query")
// 	if len(request) > 0 {
// 		request = request[1 : len(request)-1]
// 	}
// 	main_search := get_subsonic_response(writer, req, true)
// 	deezer_search := search_deezer(request)
// 	var num_songs int
// 	if len(deezer_search) < num_deezer_songs {
// 		num_songs = len(deezer_search)
// 	} else {
// 		num_songs = num_deezer_songs
// 	}
// 	if main_search.StatusCode != "failed" {
// 		main_search = extract_deezer_search(main_search, deezer_search, num_songs, 3)
// 	}
// 	var embedded EmbeddedResponse
// 	embedded.Subsonic = main_search
// 	json.NewEncoder(writer).Encode(embedded)
// }
