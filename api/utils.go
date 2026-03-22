package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func forward_headers(response *http.Response, w http.ResponseWriter) {
	for k, v := range response.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(response.StatusCode)

}

func Get_subsonic_response(writer http.ResponseWriter, req *http.Request,
	write_headers bool) SubsonicResponse {
	//Forwards to navidrome and returns the response
	upstream := navidrome_base + req.URL.Path[1:] + "?" + req.URL.RawQuery
	_, navidrome_data := Passthrough_proxy(writer, req, write_headers, upstream)
	defer navidrome_data.Close() //close when function exits
	data, err := io.ReadAll(navidrome_data)
	check_err(err)
	var response Response
	err = json.Unmarshal(data, &response)
	check_err(err)
	return response.SubsonicResponse
}

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func is_integer(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func Passthrough_proxy(w http.ResponseWriter, r *http.Request,
	write_headers bool, upstream string) (http.ResponseWriter, io.ReadCloser) {
	//Gets base data (ie default data) from navidrome and returns it
	//upstream := navidrome_base + r.URL.Path[1:] + "?" + r.URL.RawQuery
	req, err := http.NewRequest(r.Method, upstream, nil)
	check_err(err)
	//forward_headers(r, req, w)
	//Forward range headers
	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
		req.Header.Set("Range", rangeHeader)
	}
	response, err := http.DefaultClient.Do(req)
	check_err(err)
	if write_headers {
		forward_headers(response, w)
	}
	return w, response.Body
}

func copy_file(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	_, err = os.Stat(dst)
	//new := false
	var out *os.File
	if err != nil {
		out, err = os.Create(dst)
		if err != nil {
			return err
		}
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync() // flush to disk
}

// func Fetch_base_data(w http.ResponseWriter, r *http.Request,
// 	write_headers bool) (http.ResponseWriter, io.ReadCloser) {
// 	//Gets base data (ie default data) from navidrome and returns it
// 	upstream := navidrome_base + r.URL.Path[1:] + "?" + r.URL.RawQuery
// 	req, err := http.NewRequest(r.Method, upstream, nil)
// 	check_err(err)
// 	//forward_headers(r, req, w)
// 	//Forward range headers
// 	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
// 		req.Header.Set("Range", rangeHeader)
// 	}
// 	response, err := http.DefaultClient.Do(req)
// 	check_err(err)
// 	if write_headers {
// 		forward_headers(response, w)
// 	}
// 	return w, response.Body
// }
