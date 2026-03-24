package deezer

import (
	"encoding/json"
	"io"
	"net/http"
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
