package deemix

type Deemix struct {
	Arl     string `json:"arl"`
	Force   bool   `json:"force"`
	Child   int    `json:"child"`
	Url     string `json:"url"`
	Bitrate any    `json:"bitrate"`
}

type Root struct {
	Queue      map[string]QueueItem `json:"queue"`
	QueueOrder []string             `json:"queueOrder"`
}

type QueueItem struct {
	Type       string   `json:"type"`
	ID         string   `json:"id"`
	Bitrate    int      `json:"bitrate"`
	UUID       string   `json:"uuid"`
	Title      string   `json:"title"`
	Artist     string   `json:"artist"`
	Cover      string   `json:"cover"`
	Explicit   bool     `json:"explicit"`
	Size       int      `json:"size"`
	Downloaded int      `json:"downloaded"`
	Failed     int      `json:"failed"`
	Progress   int      `json:"progress"`
	Errors     []string `json:"errors"`
	Files      []File   `json:"files"`
	ExtrasPath string   `json:"extrasPath"`
	TypeTag    string   `json:"__type__"` // renamed to avoid conflict with Type
	Status     string   `json:"status"`
}

type File struct {
	Filename string   `json:"filename"`
	Data     FileData `json:"data"`
	Path     string   `json:"path"`
}

type FileData struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
