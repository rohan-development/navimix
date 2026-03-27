package types

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
	Parent              string         `json:"parent"`
	IsDir               bool           `json:"isDir"`
	Title               string         `json:"title"`
	Name                string         `json:"name"`
	Artist              string         `json:"artist"`
	ArtistID            string         `json:"artistId"`
	CoverArt            string         `json:"coverArt"`
	SongCount           int            `json:"songCount"`
	Duration            int            `json:"duration"`
	PlayCount           int            `json:"playCount"`
	Created             string         `json:"created"`
	Comment             string         `json:"comment"`
	ReplayGain          ReplayGain     `json:"replayGain"`
	BPM                 int            `json:"bpm"`
	Year                int            `json:"year"`
	Genre               string         `json:"genre"`
	Played              string         `json:"played"`
	UserRating          int            `json:"userRating"`
	Genres              []Genre        `json:"genres"`
	MediaType           string         `json:"mediaType"`
	MusicBrainzID       string         `json:"musicBrainzId"`
	IsCompilation       bool           `json:"isCompilation"`
	SortName            string         `json:"sortName"`
	OriginalReleaseDate map[string]any `json:"originalReleaseDate"`
	ReleaseDate         map[string]any `json:"releaseDate"`
	ReleaseTypes        []string       `json:"releaseTypes"`
	RecordLabels        []Label        `json:"recordLabels"`
	Moods               []string       `json:"moods"`
	Artists             []ArtistRef    `json:"artists"`
	AlbumArtists        []ArtistRef    `json:"albumArtists"`
	DisplayArtist       string         `json:"displayArtist"`
	ExplicitStatus      string         `json:"explicitStatus"`
	Version             string         `json:"version"`
	Tracks              []Song         `json:"song,omitzero"`
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
