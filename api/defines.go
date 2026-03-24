package api

import "navimix/config"

var navidrome_base string = ""
var arl string = ""

var Navidrome_base string = navidrome_base

const deezer_search_base = "https://api.deezer.com/search/"
const deezer_api_base = "https://api.deezer.com/"

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
}

type SubsonicError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type searchResult2 struct {
	Artist []Artist `json:"artist,omitempty"`
	Album  []Song   `json:"album,omitempty"`
	Song   []Song   `json:"song,omitempty"`
}

type Song struct {
	ID                 string     `json:"id"`
	Parent             string     `json:"parent"`
	IsDir              bool       `json:"isDir"`
	Title              string     `json:"title"`
	Album              string     `json:"album"`
	Name               string     `json:"name,omitempty"`
	Artist             string     `json:"artist"`
	Track              int        `json:"track"`
	Year               int        `json:"year"`
	Genre              string     `json:"genre"`
	CoverArt           string     `json:"coverArt"`
	Size               int        `json:"size"`
	ContentType        string     `json:"contentType"`
	Suffix             string     `json:"suffix"`
	Duration           int        `json:"duration,omitempty"`
	BitRate            int        `json:"bitRate"`
	Path               string     `json:"path,omitempty"`
	PlayCount          int        `json:"playCount"`
	DiscNumber         int        `json:"discNumber"`
	Created            string     `json:"created"`
	AlbumID            string     `json:"albumId"`
	ArtistID           string     `json:"artistId"`
	Type               string     `json:"type"`
	IsVideo            bool       `json:"isVideo,omitempty"`
	Played             string     `json:"played"`
	BPM                int        `json:"bpm"`
	Comment            string     `json:"comment"`
	SortName           string     `json:"sortName"`
	MediaType          string     `json:"mediaType"`
	MusicBrainzID      string     `json:"musicBrainzId"`
	ISRC               []string   `json:"isrc"`
	Genres             []Genre    `json:"genres"`
	ReplayGain         ReplayGain `json:"replayGain"`
	ChannelCount       int        `json:"channelCount"`
	SamplingRate       int        `json:"samplingRate"`
	BitDepth           int        `json:"bitDepth"`
	Moods              []Mood     `json:"moods"`
	Artists            []Artist   `json:"artists"`
	DisplayArtist      string     `json:"displayArtist"`
	AlbumArtists       []Artist   `json:"albumArtists"`
	DisplayAlbumArtist string     `json:"displayAlbumArtist"`
	Contributors       []Artist   `json:"contributors"`
	DisplayComposer    string     `json:"displayComposer"`
	ExplicitStatus     string     `json:"explicitStatus"`
}

type Album struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Artist              string         `json:"artist"`
	ArtistID            string         `json:"artistId"`
	CoverArt            string         `json:"coverArt"`
	SongCount           int            `json:"songCount"`
	Duration            int            `json:"duration"`
	PlayCount           int            `json:"playCount"`
	Created             string         `json:"created"`
	Year                int            `json:"year"`
	Genre               string         `json:"genre"`
	Played              string         `json:"played"`
	UserRating          int            `json:"userRating"`
	Genres              []Genre        `json:"genres"`
	MusicBrainzID       string         `json:"musicBrainzId"`
	IsCompilation       bool           `json:"isCompilation"`
	SortName            string         `json:"sortName"`
	DiscTitles          []string       `json:"discTitles"`
	OriginalReleaseDate map[string]any `json:"originalReleaseDate"`
	ReleaseDate         map[string]any `json:"releaseDate"`
	ReleaseTypes        []string       `json:"releaseTypes"`
	RecordLabels        []Label        `json:"recordLabels"`
	Moods               []string       `json:"moods"`
	Artists             []ArtistRef    `json:"artists"`
	DisplayArtist       string         `json:"displayArtist"`
	ExplicitStatus      string         `json:"explicitStatus"`
	Version             string         `json:"version"`
}

type ArtistRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Label struct {
	Name string `json:"name"`
}

type ISRC struct {
	Zero string `json:"0"`
}

type Genre struct {
	Name string `json:"name"`
}

type ReplayGain struct {
	TrackGain float64 `json:"trackGain"`
	AlbumGain float64 `json:"albumGain"`
	TrackPeak float64 `json:"trackPeak"`
	AlbumPeak float64 `json:"albumPeak"`
}

type Mood struct {
}

type Artist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CoverArt string `json:"coverArt,omitempty"`
	ImageUrl string `json:"artistImageUrl,omitempty"`
}

type deezer_response struct {
	Data []deezer_data `json:"data"`
}

type deezer_album struct {
	ID          int     `json:"id"`
	Name        string  `json:"title"`
	Year        string  `json:"release_date"`
	Genres      DGenres `json:"genres,omitzero"`
	CoverSmall  string  `json:"cover_small"`
	CoverMedium string  `json:"cover_medium"`
	CoverBig    string  `json:"cover_big"`
	CoverXL     string  `json:"cover_xl"`
}

type Deezer_Genre struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Type    string `json:"type"`
}

type DGenres struct {
	Data []Deezer_Genre `json:"data"`
}

type deezer_artist struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture_small"`
}

type deezer_data struct {
	ID           int           `json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Name         string        `json:"name,omitempty"`
	Cover        string        `json:"cover,omitempty"`
	CoverSmall   string        `json:"cover_small,omitempty"`
	CoverMedium  string        `json:"cover_medium,omitempty"`
	CoverBig     string        `json:"cover_big,omitempty"`
	ArtistSmall  string        `json:"picture_small,omitempty"`
	ArtistMedium string        `json:"picture_medium,omitempty"`
	ArtistBig    string        `json:"picture_big,omitempty"`
	ISRC         string        `json:"isrc,omitempty"`
	Link         string        `json:"link,omitempty"`
	Duration     int           `json:"duration,omitempty"`
	Artist       deezer_artist `json:"artist,omitzero"`
	Album        deezer_album  `json:"album,omitzero"`
	Type         string        `json:"type,omitzero"`
}

func Loadconfig(conf *config.Config) {
	if conf.NavidromeAddress != "" {
		navidrome_base = conf.NavidromeAddress
		Navidrome_base = navidrome_base
	}
	arl = conf.DeezerARL
}
