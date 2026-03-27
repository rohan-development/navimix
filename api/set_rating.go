package api

import (
	"io"
	"navimix/deemix"
	"net/http"
	"os"
	"path/filepath"
)

func SetRating(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if is_integer(id) {
		//If not in library, add it
		client := deemix.Login(arl, deemix_per)
		deemix.AddToQueue(id, deemix_per, client)
	}
	//Forward request regardless
	upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
	writer, response := Passthrough_proxy(writer, req, true, upstream)
	defer response.Close()
	io.Copy(writer, response)
	_, err := os.Stat(filepath.Join("music", id+".mp3"))
	if err == nil {
		os.Remove(filepath.Join("music", id+".mp3"))

	}
}
