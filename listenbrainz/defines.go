package listenbrainz

import "navimix/config"

var listenbrainz_base string = "https://api.listenbrainz.org/1/"

func Loadconfig(conf *config.Config) {
	listenbrainz_base = conf.ListenbrainzAddress
}

type ListenSubmission struct {
	ListenType string    `json:"listen_type"`
	Payload    []Payload `json:"payload"`
}

type Payload struct {
	ListenedAt    int64         `json:"listened_at,omitempty"`
	TrackMetadata TrackMetadata `json:"track_metadata"`
}

type TrackMetadata struct {
	ArtistName     string         `json:"artist_name"`
	TrackName      string         `json:"track_name"`
	ReleaseName    string         `json:"release_name"`
	AdditionalInfo AdditionalInfo `json:"additional_info"`
	MBIDMapping    MBIDMapping    `json:"mbid_mapping"`
}

type AdditionalInfo struct {
	ArtistMBIDs             []string `json:"artist_mbids"`
	ArtistNames             []string `json:"artist_names"`
	DurationMS              int      `json:"duration_ms"`
	SubmissionClient        string   `json:"submission_client"`
	SubmissionClientVersion string   `json:"submission_client_version"`
}

type MBIDMapping struct {
	ArtistMBIDs   interface{} `json:"artist_mbids"`   // null in JSON
	Artists       interface{} `json:"artists"`        // null in JSON
	RecordingMBID string      `json:"recording_mbid"` // empty string
	ReleaseMBID   string      `json:"release_mbid"`   // empty string
}
