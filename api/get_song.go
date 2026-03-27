package api

import (
	"encoding/json"
	"io"
	"navimix/db"
	"navimix/deezer"
	"navimix/types"
	"net/http"
	"strconv"
)

func GetSong(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if !is_integer(id) {
		//in library, forward
		upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		writer, response := Passthrough_proxy(writer, req, true, upstream)
		defer response.Close()
		io.Copy(writer, response)
	} else {
		var track types.Song
		deezer_search, err := db.GetTrack(id)
		if err != nil {
			deezer_search = deezer.GetTrack(id)
			db.AddTrack(deezer_search)
		}
		track = populate_song(track, deezer_search)
		if track.Album == "" && track.Title == "" {
			upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
			writer, response := Passthrough_proxy(writer, req, true, upstream)
			defer response.Close()
			io.Copy(writer, response)
			return
		}
		album := deezer.GetAlbum(track.AlbumID)
		track.Genre = album.Genres.Data[0].Name
		track.Year, _ = strconv.Atoi(album.Year[0:4])
		var embedded EmbeddedResponse
		embedded.Subsonic = Get_subsonic_response(writer, req, true)
		embedded.Subsonic.StatusCode = "ok"
		embedded.Subsonic.Error = SubsonicError{}
		embedded.Subsonic.Song = track
		//embedded.Subsonic.StatusCode = req.URL.Path[6:]
		json.NewEncoder(writer).Encode(embedded)
	}
}
