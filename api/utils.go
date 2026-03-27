package api

import (
	"encoding/json"
	"io"
	"log"
	"navimix/deezer"
	"navimix/types"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func forward_headers(response *http.Response, w http.ResponseWriter) {
	for k, v := range response.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(response.StatusCode)

}

func Get_subsonic_response(writer http.ResponseWriter, req *http.Request,
	write_headers bool) SubsonicResponse {
	//Forwards to navidrome and returns the response
	upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
	_, navidrome_data := Passthrough_proxy(writer, req, write_headers, upstream)
	defer navidrome_data.Close() //close when function exits
	data, err := io.ReadAll(navidrome_data)
	check_err(err)
	var response Response
	err = json.Unmarshal(data, &response)
	check_err(err)
	return response.SubsonicResponse
}

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func is_integer(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func Passthrough_proxy(w http.ResponseWriter, r *http.Request,
	write_headers bool, upstream string) (http.ResponseWriter, io.ReadCloser) {
	//Gets base data (ie default data) from navidrome and returns it
	//upstream := navidrome_base + r.URL.Path[1:] + "?" + r.URL.RawQuery
	req, err := http.NewRequest(r.Method, upstream, nil)
	check_err(err)
	//forward_headers(r, req, w)
	//Forward range headers
	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
		req.Header.Set("Range", rangeHeader)
	}
	response, err := http.DefaultClient.Do(req)
	check_err(err)
	if write_headers {
		forward_headers(response, w)
	}
	return w, response.Body
}

func copy_file(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	_, err = os.Stat(dst)
	//new := false
	var out *os.File
	if err != nil {
		out, err = os.Create(dst)
		if err != nil {
			return err
		}
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync() // flush to disk
}

func populate_song(add_song types.Song, deezer_search deezer.Data) types.Song {
	add_song.ID = strconv.Itoa(deezer_search.ID)
	add_song.Title = deezer_search.Title
	add_song.Album = deezer_search.Album.Name
	add_song.Artist = deezer_search.Artist.Name
	add_song.AlbumID = strconv.Itoa(deezer_search.Album.ID)
	add_song.Parent = add_song.AlbumID
	add_song.BitRate = 128
	add_song.ContentType = "audio/mp3"
	add_song.Suffix = "mp3"
	// album := query_deezer_api("album/" + add_song.AlbumID)
	// add_song.Year, _ = strconv.Atoi(album.Year[0:4])
	// add_song.Genre = album.Genres.Data[0].Name
	add_song.Duration = deezer_search.Duration
	add_song.SortName = strings.ToLower(add_song.Title)
	add_song.Type = "music"
	add_song.MediaType = "song"
	add_song.DisplayArtist = add_song.Artist
	return add_song
}

func populate_album(add_album types.Album, deezer_search deezer.Data) types.Album {
	add_album.ID = strconv.Itoa(deezer_search.ID)
	add_album.Name = deezer_search.Title
	add_album.Title = deezer_search.Title
	add_album.Artist = deezer_search.Artist.Name
	add_album.Parent = strconv.Itoa(deezer_search.Artist.ID)
	add_album.MediaType = "album"
	add_album.DisplayArtist = add_album.Artist
	// album := query_deezer_api("album/" + add_album.ID)
	// add_album.Year, _ = strconv.Atoi(album.Year[0:4])
	// add_album.Genre = album.Genres.Data[0].Name
	return add_album
}
