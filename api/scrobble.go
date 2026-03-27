package api

import (
	"encoding/json"
	"io"
	"navimix/deezer"
	"navimix/listenbrainz"
	"net/http"
)

func Scrobble(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if !is_integer(id) {
		//in library, forward
		upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		writer, response := Passthrough_proxy(writer, req, true, upstream)
		defer response.Close()
		io.Copy(writer, response)
		return
	} else if listenbrainz_enabled && req.URL.Query().Get("u") == listenbrainz_user {
		track := deezer.GetTrack(id)
		switch req.URL.Query().Get("submission") {
		case "true", "":
			listenbrainz.Scrobble(track.Artist.Name, track.Album.Name,
				track.Title, req.URL.Query().Get("time"), listenbrainz_api,
				track.Duration*1000, true)
		case "false":
			listenbrainz.Scrobble(track.Artist.Name, track.Album.Name,
				track.Title, "", listenbrainz_api, track.Duration*1000, false)
		}
	}
	var embedded EmbeddedResponse
	embedded.Subsonic = Get_subsonic_response(writer, req, true)
	json.NewEncoder(writer).Encode(embedded)
}
