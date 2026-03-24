package deezer

const deezer_search_base = "https://api.deezer.com/search/"
const deezer_api_base = "https://api.deezer.com/"

type Response struct {
	Data []Data `json:"data"`
}

type Album struct {
	ID          int      `json:"id"`
	Name        string   `json:"title"`
	Year        string   `json:"release_date"`
	Artist      Artist   `json:"artist"`
	Duration    int      `json:"duration"`
	Genres      Genres   `json:"genres,omitzero"`
	CoverSmall  string   `json:"cover_small"`
	CoverMedium string   `json:"cover_medium"`
	CoverBig    string   `json:"cover_big"`
	CoverXL     string   `json:"cover_xl"`
	Tracks      Response `json:"tracks,omitzero"`
}

type AlbumRef struct {
	ID          int    `json:"id"`
	Name        string `json:"title"`
	Year        string `json:"release_date"`
	Artist      Artist `json:"artist"`
	Duration    int    `json:"duration"`
	Genres      Genres `json:"genres,omitzero"`
	CoverSmall  string `json:"cover_small"`
	CoverMedium string `json:"cover_medium"`
	CoverBig    string `json:"cover_big"`
	CoverXL     string `json:"cover_xl"`
	//	Tracks      Data   `json:"tracks,omitzero"`
}

type Genre struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Type    string `json:"type"`
}

type Genres struct {
	Data []Genre `json:"data"`
}

type Artist struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture_small"`
}

type Data struct {
	ID           int      `json:"id,omitempty"`
	Title        string   `json:"title,omitempty"`
	Name         string   `json:"name,omitempty"`
	Cover        string   `json:"cover,omitempty"`
	CoverSmall   string   `json:"cover_small,omitempty"`
	CoverMedium  string   `json:"cover_medium,omitempty"`
	CoverBig     string   `json:"cover_big,omitempty"`
	ArtistSmall  string   `json:"picture_small,omitempty"`
	ArtistMedium string   `json:"picture_medium,omitempty"`
	ArtistBig    string   `json:"picture_big,omitempty"`
	ISRC         string   `json:"isrc,omitempty"`
	Link         string   `json:"link,omitempty"`
	Duration     int      `json:"duration,omitempty"`
	Artist       Artist   `json:"artist,omitzero"`
	Album        AlbumRef `json:"album,omitzero"`
	Type         string   `json:"type,omitzero"`
}
