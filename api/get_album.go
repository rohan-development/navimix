package api

import (
	"encoding/json"
	"io"
	"navimix/db"
	"navimix/deezer"
	"navimix/types"
	"net/http"
	"net/url"
	"strconv"
)

func GetAlbum(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	var album types.Album
	if !is_integer(id) {
		//in library, forward
		// upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		// writer, response := Passthrough_proxy(writer, req, true, upstream)
		// defer response.Close()
		// io.Copy(writer, response)

		local_album := Get_subsonic_response(writer, req, false).Album
		deezer_search1 := deezer.Search(local_album.Name+" "+
			local_album.Artists[0].Name, "album")
		//fmt.Fprint(writer, local_album.Name+" "+local_album.Artists[0].Name)
		//var deezer_album deezer.AlbumRef
		if len(deezer_search1) > 0 {
			deezer_search, err := db.GetAlbum(strconv.Itoa(deezer_search1[0].ID))
			deezer_search = deezer.GetAlbum(strconv.Itoa(deezer_search1[0].ID))
			if err != nil {
				db.AddAlbum(deezer_search)
			}
			album = deezer_to_album(deezer_search)
			for i := 0; i < album.SongCount; i += 1 {
				var add_song types.Song
				add_song.Track = i + 1
				overide := false
				for j := 0; j < len(local_album.Tracks); j += 1 {
					if local_album.Tracks[j].Track == add_song.Track {
						album.Tracks = append(album.Tracks, local_album.Tracks[j])
						overide = true
						break
					}
				}

				if !overide {
					add_song = populate_song(add_song, deezer_search.Tracks.Data[i])
					add_song.Parent = id
					add_song.AlbumID = id
					if add_song.DiscNumber == 0 {
						add_song.DiscNumber = 1
					}
					album.Tracks = append(album.Tracks, add_song)

				}

			}

		}
		local_album.Tracks = album.Tracks
		album = local_album

	} else {
		//var album types.Album
		deezer_search, err := db.GetAlbum(id)
		deezer_search = deezer.GetAlbum(id)
		if err != nil {
			db.AddAlbum(deezer_search)
		}
		album = deezer_to_album(deezer_search)
		if album.Name == "" {
			upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
			writer, response := Passthrough_proxy(writer, req, true, upstream)
			defer response.Close()
			io.Copy(writer, response)
			return
		}
		alb_url := build_search_URL(req, album, "search2.view")
		alb_req, err := http.NewRequest("GET", alb_url, nil)
		check_err(err)
		local_album_search := Get_subsonic_response(writer, alb_req, false).SearchResult2.
			Album
		var local_album types.Album
		if len(local_album_search) != 0 {
			//local_album_id := local_album_search[0].ID
			alb_url = build_search_URL(req, local_album_search[0], "getAlbum.view")
			//fmt.Fprint(writer, alb_url)
			alb_req, _ = http.NewRequest("GET", alb_url, nil)
			local_album = Get_subsonic_response(writer, alb_req, false).Album
		}
		for i := 0; i < album.SongCount; i += 1 {
			var add_song types.Song
			add_song.Track = i + 1
			overide := false

			for j := 0; j < len(local_album.Tracks); j += 1 {
				if local_album.Tracks[j].Track == add_song.Track {
					local_album.Tracks[j].AlbumID = id
					local_album.Tracks[j].Parent = id
					album.Tracks = append(album.Tracks, local_album.Tracks[j])
					overide = true
					break
				}
			}

			if !overide {
				add_song = populate_song(add_song, deezer_search.Tracks.Data[i])
				if add_song.DiscNumber == 0 {
					add_song.DiscNumber = 1
				}
				album.Tracks = append(album.Tracks, add_song)
			}

		}
	}
	var embedded EmbeddedResponse
	embedded.Subsonic = Get_subsonic_response(writer, req, true)
	embedded.Subsonic.StatusCode = "ok"
	embedded.Subsonic.Error = SubsonicError{}
	embedded.Subsonic.Album = album
	//embedded.Subsonic.StatusCode = req.URL.Path[6:]
	json.NewEncoder(writer).Encode(embedded)

}

func build_search_URL(req *http.Request, album types.Album, method string) string {
	params := url.Values{}
	params.Set("u", req.URL.Query().Get("u"))
	params.Set("p", req.URL.Query().Get("p"))
	params.Set("t", req.URL.Query().Get("t"))
	params.Set("s", req.URL.Query().Get("s"))
	params.Set("v", req.URL.Query().Get("v"))
	params.Set("c", req.URL.Query().Get("c"))

	params.Set("f", "json")
	params.Set("artistCount", "0")
	params.Set("albumCount", "1")
	params.Set("songCount", "0")
	if method == "search2.view" {
		params.Set("query", album.Name+" "+album.Artist)
	} else if method == "getAlbum.view" {
		params.Set("id", album.ID)
	}

	return (navidrome_base + "rest/" + method + "?" + params.Encode())
}

func deezer_to_album(deezer_search deezer.Album) types.Album {
	var album types.Album
	album.ID = strconv.Itoa(deezer_search.ID)
	album.Name = deezer_search.Name
	album.Artist = deezer_search.Artist.Name
	album.Parent = strconv.Itoa(deezer_search.Artist.ID)
	album.Duration = deezer_search.Duration
	album.SongCount = len(deezer_search.Tracks.Data)
	album.MediaType = "album"
	album.DisplayArtist = album.Artist
	album.Year, _ = strconv.Atoi(deezer_search.Year[0:4])
	album.Genre = deezer_search.Genres.Data[0].Name
	album.CoverArt = album.ID
	return album
}
