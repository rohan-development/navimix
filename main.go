package main

import (
	"io"
	"net/http"
)

func main() {
	init_db()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, r *http.Request) {
	//Get rid of /rest/
	api_call := r.URL.Path[6:]
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"search2":     handleSearch,
		"stream":      handleStream,
		"download":    handleDownload,
		"getCoverArt": handleCoverArt,
		"getLyrics":   handleLyrics,
		"scrobble":    handleScrobble,
		"getSong":     handleInfo,
	}
	routing, special_api := routes[api_call]
	if special_api {
		routing(writer, r)
		//writer, response := fetch_base_data(writer, r)
		//defer response.Close()
		//io.Copy(writer, response)
	} else {
		writer, response := fetch_base_data(writer, r, true)
		defer response.Close()
		io.Copy(writer, response)
	}
}
