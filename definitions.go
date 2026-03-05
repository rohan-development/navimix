package main

const deemix_base = "http://localhost:6595"
const navidrome_base = "http://localhost:4533/"

const deezer_search_base = "https://api.deezer.com/search?q="

type Response struct {
	SubsonicResponse SubsonicResponse `json:"subsonic-response"`
}

type SubsonicResponse struct {
	StatusCode    string        `json:"status"`
	Version       string        `json:"version"`
	Type          string        `json:"type"`
	ServerVersion string        `json:"serverVersion"`
	OpenSubsonic  bool          `json:"openSubsonic"`
	SearchResult2 searchResult2 `json:"searchResult2,omitempty"`
	Error         SubsonicError `json:"error,omitempty"`
}

type SubsonicError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type searchResult2 struct {
	Song []Song `json:"song,omitempty"`
}

type Song struct {
	ID                 string     `json:"id"`
	Parent             string     `json:"parent"`
	IsDir              bool       `json:"isDir"`
	Title              string     `json:"title"`
	Album              string     `json:"album"`
	Artist             string     `json:"artist"`
	Track              int        `json:"track"`
	Year               int        `json:"year"`
	Genre              string     `json:"genre"`
	CoverArt           string     `json:"coverArt"`
	Size               int        `json:"size"`
	ContentType        string     `json:"contentType"`
	Suffix             string     `json:"suffix"`
	Duration           int        `json:"duration"`
	BitRate            int        `json:"bitRate"`
	Path               string     `json:"path"`
	PlayCount          int        `json:"playCount"`
	DiscNumber         int        `json:"discNumber"`
	Created            string     `json:"created"`
	AlbumID            string     `json:"albumId"`
	ArtistID           string     `json:"artistId"`
	Type               string     `json:"type"`
	IsVideo            bool       `json:"isVideo"`
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
	ID   string `json:"id"`
	Name string `json:"name"`
}

type deezer_response struct {
	Data []deezer_data `json:"data"`
}

type deezer_album struct {
	ID      int    `json:"id"`
	Name    string `json:"title"`
	Picture string `json:"cover_small"`
}

type deezer_artist struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture_small"`
}

type deezer_data struct {
	ID       int           `json:"id"`
	Title    string        `json:"title"`
	ISRC     string        `json:"isrc"`
	Link     string        `json:"link"`
	Duration int           `json:"duration"`
	Artist   deezer_artist `json:"artist"`
	Album    deezer_album  `json:"album"`
}
