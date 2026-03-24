package listenbrainz

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

func Scrobble(artist, album, track, time, api string, submission bool) {
	var payload []byte
	//unix_time := strconv.Itoa(time)

	var scrobble ListenSubmission
	var tmp_payload Payload
	scrobble.Payload = append(scrobble.Payload, tmp_payload)
	if submission {
		scrobble.ListenType = "single"
		if time != "" {
			scrobble.Payload[0].ListenedAt, _ = strconv.ParseInt(time, 10, 64)
			scrobble.Payload[0].ListenedAt /= 1000 //Navidrome gives ms, convert to s
		}
	} else {
		scrobble.ListenType = "playing_now"
	}
	scrobble.Payload[0].TrackMetadata.ReleaseName = album
	scrobble.Payload[0].TrackMetadata.ArtistName = artist
	scrobble.Payload[0].TrackMetadata.TrackName = track
	payload, err := json.Marshal(scrobble)
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST",
		listenbrainz_base+"submit-listens",
		bytes.NewBuffer(payload),
	)
	req.Header.Set("Authorization", "Token "+api)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)
}
