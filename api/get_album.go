package api

import (
	"encoding/json"
	"io"
	"navimix/deezer"
	"navimix/types"
	"net/http"
	"strconv"
)

func GetAlbum(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if !is_integer(id) {
		//in library, forward
		upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		writer, response := Passthrough_proxy(writer, req, true, upstream)
		defer response.Close()
		io.Copy(writer, response)
	} else {
		var album types.Album
		deezer_search := deezer.GetAlbum(id)
		album.ID = strconv.Itoa(deezer_search.ID)
		album.Name = deezer_search.Name
		if album.Name == "" {
			upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
			writer, response := Passthrough_proxy(writer, req, true, upstream)
			defer response.Close()
			io.Copy(writer, response)
			return
		}
		album.Artist = deezer_search.Artist.Name
		album.Parent = strconv.Itoa(deezer_search.Artist.ID)
		album.Duration = deezer_search.Duration
		album.SongCount = len(deezer_search.Tracks.Data)
		album.MediaType = "album"
		album.DisplayArtist = album.Artist
		album.Year, _ = strconv.Atoi(deezer_search.Year[0:4])
		album.Genre = deezer_search.Genres.Data[0].Name
		album.CoverArt = album.ID
		for i := 0; i < album.SongCount; i += 1 {
			var add_song types.Song
			add_song.Track = i + 1
			album.Tracks = append(album.Tracks,
				populate_song(add_song, deezer_search.Tracks.Data[i]))
		}
		var embedded EmbeddedResponse
		embedded.Subsonic = Get_subsonic_response(writer, req, true)
		embedded.Subsonic.StatusCode = "ok"
		embedded.Subsonic.Error = SubsonicError{}
		embedded.Subsonic.Album = album
		//embedded.Subsonic.StatusCode = req.URL.Path[6:]
		json.NewEncoder(writer).Encode(embedded)
	}
}
