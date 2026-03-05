package main

import (
	"encoding/json"
	"net/http"
)

const num_deezer_songs = 3

func handleSearch(writer http.ResponseWriter, req *http.Request) {
	request := req.URL.Query().Get("query")
	request = request[1 : len(request)-1]
	main_search := get_subsonic_response(writer, req, true)
	deezer_search := search_deezer(request)
	var num_songs int
	if len(deezer_search) < num_deezer_songs {
		num_songs = len(deezer_search)
	} else {
		num_songs = num_deezer_songs
	}
	if main_search.StatusCode != "failed" {
		main_search = extract_deezer_search(main_search, deezer_search, num_songs)
	}
	json.NewEncoder(writer).Encode(main_search)
}

func handleStream(w http.ResponseWriter, r *http.Request) {

}

func handleDownload(w http.ResponseWriter, r *http.Request) {

}

func handleCoverArt(w http.ResponseWriter, r *http.Request) {

}

func handleLyrics(http.ResponseWriter, *http.Request) {

}

func handleScrobble(http.ResponseWriter, *http.Request) {

}

func handleInfo(http.ResponseWriter, *http.Request) {

}
