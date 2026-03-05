package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func song_in_search(main_search SubsonicResponse, add_song Song) bool {
	for j := 0; j < len(main_search.SearchResult2.Song); j += 1 {
		if add_song.Title == main_search.SearchResult2.Song[j].Title &&
			add_song.Artist == main_search.SearchResult2.Song[j].Artist {
			return true
		}
	}
	return false
}

func extract_deezer_search(main_search SubsonicResponse,
	deezer_search []deezer_data, num_songs int) SubsonicResponse {
	var add_song Song
	delta := 0
	for i := 0; i < num_songs; i += 1 {
		add_song.ID = strconv.Itoa(deezer_search[i+delta].ID)
		add_song.Title = deezer_search[i+delta].Title
		add_song.Album = deezer_search[i+delta].Album.Name
		add_song.Artist = deezer_search[i+delta].Artist.Name
		//song_in_search := false
		// for j := 0; j < len(main_search.SearchResult2.Song); j += 1 {
		// 	if add_song.Title == main_search.SearchResult2.Song[j].Title &&
		// 		add_song.Artist == main_search.SearchResult2.Song[j].Artist {
		// 		song_in_search = true
		// 		break
		// 	}
		// }
		if song_in_search(main_search, add_song) {
			delta += 1
			i -= 1
			continue
		}
		// add_song.ISRC = deezer_search[i].ISRC
		add_song.Duration = deezer_search[i+delta].Duration
		add_song.SortName = strings.ToLower(add_song.Title)
		add_song.Type = "song"
		add_song.DisplayArtist = add_song.Artist
		//add_song.Artists[0].Name = add_song.Artist
		main_search.SearchResult2.Song = append(main_search.SearchResult2.Song,
			add_song)
		if !song_exists_in_db(deezer_search[i+delta].ID) {
			add_song_to_database(add_song)
		}
	}
	return main_search
}

func forward_headers(response *http.Response, w http.ResponseWriter) {
	for k, v := range response.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(response.StatusCode)

}

func fetch_base_data(w http.ResponseWriter, r *http.Request,
	write_headers bool) (http.ResponseWriter, io.ReadCloser) {
	//Gets base data (ie default data) from navidrome and returns it
	upstream := navidrome_base + r.URL.Path[1:] + "?" + r.URL.RawQuery
	req, err := http.NewRequest(r.Method, upstream, nil)
	check_err(err)
	//forward_headers(r, req, w)
	//Forward range headers
	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
		req.Header.Set("Range", rangeHeader)
	}
	response, err := http.DefaultClient.Do(req)
	check_err(err)
	if write_headers {
		forward_headers(response, w)
	}
	return w, response.Body
}

func get_subsonic_response(writer http.ResponseWriter, req *http.Request,
	write_headers bool) SubsonicResponse {
	//Forwards to navidrome and returns the response
	_, navidrome_data := fetch_base_data(writer, req, write_headers)
	defer navidrome_data.Close() //close when function exits
	data, err := io.ReadAll(navidrome_data)
	check_err(err)
	var response Response
	err = json.Unmarshal(data, &response)
	check_err(err)
	return response.SubsonicResponse
}

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

// func http_get_string(url string) (string, error) {
// 	resp, err := http.Get(url)
// 	//fmt.Print(err)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	resp.Body.Close()
// 	return string(body), err
// }

// func navidrome_query(endpoint, format string, arguments ...string) (string, error) {
// 	url := navidrome_base + "/rest/" + endpoint + "?&" + navidrome_info + "&f=" + format
// 	for _, s := range arguments {
// 		url += "&" + s
// 	}
// 	return http_get_string(url)
// }
