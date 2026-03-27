package api

import (
	"navimix/config"
	"navimix/types"
)

var navidrome_base string = ""
var deemix_tmp string = "http://localhost:6596/"
var deemix_per string = "http://localhost:6595/"
var arl string = ""
var listenbrainz_api string = ""
var listenbrainz_enabled bool = false
var listenbrainz_user string = ""

var Navidrome_base string = navidrome_base

type Response struct {
	SubsonicResponse SubsonicResponse `json:"subsonic-response"`
}
type EmbeddedResponse struct {
	Subsonic SubsonicResponse `json:"subsonic-response"`
}
type SubsonicResponse struct {
	StatusCode    string        `json:"status"`
	Version       string        `json:"version"`
	Type          string        `json:"type"`
	ServerVersion string        `json:"serverVersion"`
	OpenSubsonic  bool          `json:"openSubsonic"`
	SearchResult2 searchResult2 `json:"searchResult2,omitzero"`
	SearchResult3 searchResult2 `json:"searchResult3,omitzero"`
	Error         SubsonicError `json:"error,omitzero"`
	Song          types.Song    `json:"song,omitzero"`
	Album         types.Album   `json:"album,omitzero"`
	Artist        types.Artist  `json:"artist,omitzero"`
}

type SubsonicError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type searchResult2 struct {
	Artist []types.Artist `json:"artist,omitempty"`
	Album  []types.Album  `json:"album,omitempty"`
	Song   []types.Song   `json:"song,omitempty"`
}

func Loadconfig(conf *config.Config) {
	if conf.NavidromeAddress != "" {
		navidrome_base = conf.NavidromeAddress
		Navidrome_base = navidrome_base
	}
	arl = conf.DeezerARL
	if conf.DeemixTmp != "" {
		deemix_tmp = conf.DeemixTmp
	}
	if conf.DeemixPersistent != "" {
		deemix_per = conf.DeemixPersistent
	}
	listenbrainz_api = conf.ListenbrainzAuth
	listenbrainz_enabled = conf.ListenbrainzEnabled
	listenbrainz_user = conf.ListenbrainzUser
}
