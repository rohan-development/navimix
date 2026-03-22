package api

import (
	"io"
	"navimix/deemix"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const SLEEP_TIME = 10 //ms

func Stream(writer http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if !is_integer(id) {
		//in library, forward
		upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
		writer, response := Passthrough_proxy(writer, req, true, upstream)
		defer response.Close()
		io.Copy(writer, response)
	} else {
		//download := false
		//if req.URL.Path[6:] == "download.view" {
		//download = true
		//}
		//file, err := os.Open("track.mp3")
		//check_err(err)
		//defer file.Close()
		writer.Header().Set("Content-Type", "audio/mpeg")
		var file string
		file = filepath.Join("music", id+".mp3")
		_, err := os.Stat(file)
		new := false
		if err != nil {
			new = true
			deemix.Login(arl)
			deemix.AddToQueue(id)
			file = filepath.Join("downloads", id+".mp3")
			for !deemix.IsDone(id) {
				time.Sleep(SLEEP_TIME * time.Millisecond)
			}
		}
		http.ServeFile(writer, req, file)
		if new {
			copy_file(file, filepath.Join("music", id+".mp3"))
			os.Remove(file)
		}
	}

}
