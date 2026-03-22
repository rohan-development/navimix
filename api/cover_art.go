package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func CoverArt(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if !is_integer(id) {
		//in library, forward
		upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		writer, response := Passthrough_proxy(writer, req, true, upstream)
		defer response.Close()
		io.Copy(writer, response)
	} else {
		var art_link string
		album := query_deezer_api("album/" + id)
		quality, err := strconv.Atoi(req.URL.Query().Get("size"))
		if err == nil {
			if quality < 60 {
				art_link = album.CoverSmall
			} else if quality < 250 {
				art_link = album.CoverMedium
			} else if quality < 500 {
				art_link = album.CoverBig
			} else {
				art_link = album.CoverXL
			}
		} else {
			art_link = album.CoverBig
		}
		//fmt.Fprint(writer, art_link)
		//Passthrough_proxy(writer, req, true, art_link)
		if art_link == "" {
			upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
			writer, response := Passthrough_proxy(writer, req, true, upstream)
			defer response.Close()
			io.Copy(writer, response)
		} else {
			writer, response := Passthrough_proxy(writer, req, true, art_link)
			defer response.Close()
			io.Copy(writer, response)
		}

	}
	//deezer_info = query_deezer_api(id).Album.Cover

}

func query_deezer_api(query string) deezer_album {
	url := deezer_api_base + query
	// if attribute == "all" {
	// 	url = deezer_search_base + url.QueryEscape(":"+query)
	// }

	response, err := http.Get(url)
	check_err(err)
	data, err := io.ReadAll(response.Body)
	var link deezer_album
	err = json.Unmarshal(data, &link)
	check_err(err)
	return link
}
