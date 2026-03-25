package api

import (
	"io"
	"net/http"
)

func Star(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if !is_integer(id) {
		//in library, forward
		upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		writer, response := Passthrough_proxy(writer, req, true, upstream)
		defer response.Close()
		io.Copy(writer, response)
	} else {

	}
}
