package main

import (
	"fmt"
	"io"
	"navimix/api"
	"navimix/auth"
	"navimix/config"
	_ "navimix/db"
	"navimix/listenbrainz"
	"net/http"
	"strings"
)

func main() {
	config := config.Load()
	if config.Port == "" {
		config.Port = "4534"
	}
	api.Loadconfig(config)
	listenbrainz.Loadconfig(config)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+config.Port, nil)
}

func handler(writer http.ResponseWriter, r *http.Request) {
	//Get rid of /rest/
	//api_call := r.URL.Path[6:]
	api_call := strings.TrimPrefix(r.URL.Path, "/rest/")
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"search2.view":     api.Search,
		"search3.view":     api.Search,
		"stream.view":      api.Stream,
		"download.view":    api.Stream,
		"getCoverArt.view": api.CoverArt,
		"getSong.view":     api.GetSong,
		"getAlbum.view":    api.GetAlbum,
		"scrobble.view":    api.Scrobble,
		"setRating.view":   api.SetRating,
		"star.view":        api.SetRating,
		//"getLyrics":        api.Lyrics,

	}
	routing, special_api := routes[api_call]
	if special_api {
		temp_req := r
		if auth.Check(writer, temp_req) {
			routing(writer, r)
		} else {
			fmt.Fprint(writer, "Auth Failed")
		}
		//writer, response := fetch_base_data(writer, r)
		//defer response.Close()
		//io.Copy(writer, response)
	} else {
		upstream := api.Navidrome_base + r.URL.Path[1:] + "?" + r.URL.RawQuery
		writer, response := api.Passthrough_proxy(writer, r, true, upstream)
		defer response.Close()
		io.Copy(writer, response)

	}
}
