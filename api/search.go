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

const num_deezer_songs = 3

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
	deezer_search := search_deezer(request)
	var num_songs int
	if len(deezer_search) < num_deezer_songs {
		num_songs = len(deezer_search)
	} else {
		num_songs = num_deezer_songs
	}
	if main_search.StatusCode != "failed" {
		main_search = extract_deezer_elements(main_search, deezer_search, num_songs,
			search_version)
	}
	var embedded EmbeddedResponse
	embedded.Subsonic = main_search
	//embedded.Subsonic.StatusCode = req.URL.Path[6:]
	json.NewEncoder(writer).Encode(embedded)
}

func search_deezer(query string) []deezer_data {
	url := deezer_search_base + url.QueryEscape(query)
	// if attribute == "all" {
	// 	url = deezer_search_base + url.QueryEscape(":"+query)
	// }
	response, err := http.Get(url)
	check_err(err)
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

func extract_deezer_elements(main_search SubsonicResponse,
	deezer_search []deezer_data, num_songs, search_version int) SubsonicResponse {
	var add_song Song
	delta := 0
	for i := 0; i < num_songs; i += 1 {
		add_song.ID = strconv.Itoa(deezer_search[i+delta].ID)
		add_song.Title = deezer_search[i+delta].Title
		add_song.Album = deezer_search[i+delta].Album.Name
		add_song.Artist = deezer_search[i+delta].Artist.Name
		add_song.AlbumID = strconv.Itoa(deezer_search[i+delta].Album.ID)
		add_song.Parent = add_song.AlbumID
		add_song.BitRate = 128
		add_song.ContentType = "audio/mp3"
		add_song.Suffix = "mp3"
		add_song.Year, _ = strconv.Atoi(query_deezer_api("album/" + add_song.AlbumID).Year[0:4])
		// add_song.ISRC = deezer_search[i].ISRC
		add_song.Duration = deezer_search[i+delta].Duration
		add_song.SortName = strings.ToLower(add_song.Title)
		add_song.Type = "song"
		add_song.DisplayArtist = add_song.Artist
		//add_song.Artists[0].Name = add_song.Artist
		// if !internaldb.Element_exists(deezer_search[i+delta].ID) { //add track to db
		// 	internaldb.Add_element(deezer_search[i+delta].ID, add_song)
		// }
		// if !internaldb.Element_exists(deezer_search[i+delta].Album.ID) { //add album to db
		// 	internaldb.Add_element(deezer_search[i+delta].ID, deezer_search[i+delta].Album)
		// }
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

		// if !song_exists_in_db(deezer_search[i+delta].ID) {
		// 	add_song_to_database(add_song)
		// }
	}
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
