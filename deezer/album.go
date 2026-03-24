package deezer

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAlbum(query string) Album {
	url := deezer_api_base + "album/" + query
	response, err := http.Get(url)
	check_err(err)
	data, err := io.ReadAll(response.Body)
	var link Album
	err = json.Unmarshal(data, &link)
	check_err(err)
	return link
}
