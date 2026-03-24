package deezer

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetArtist(query string) Artist {
	url := deezer_api_base + "artist/" + query
	response, err := http.Get(url)
	check_err(err)
	data, err := io.ReadAll(response.Body)
	var link Artist
	err = json.Unmarshal(data, &link)
	check_err(err)
	return link
}
