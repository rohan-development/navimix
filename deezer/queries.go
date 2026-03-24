package deezer

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func search_deezer(query, attribute string) []Data {
	url := deezer_search_base + attribute +
		"?q=" + url.QueryEscape(query)
	// if attribute == "all" {
	// 	url = deezer_search_base + url.QueryEscape(":"+query)
	// }
	response, err := http.Get(url)
	check_err(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	var link Response
	err = json.Unmarshal(data, &link)
	check_err(err)
	return link.Data
}
