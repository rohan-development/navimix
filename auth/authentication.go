package auth

import (
	"navimix/api"
	"net/http"
)

func Check(writer http.ResponseWriter, r *http.Request) bool {
	req := clone_request(r)
	req.URL.Path = "/rest/ping.view"
	q := req.URL.Query()
	q.Set("f", "json")
	req.URL.RawQuery = q.Encode()
	response := api.Get_subsonic_response(writer, req, false)
	if response.StatusCode == "ok" {
		return true
	}
	return false
}

func clone_request(req *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := *req

	// copy URL (so we can change Path/RawQuery safely)
	urlCopy := *req.URL
	r2.URL = &urlCopy

	// copy query values
	q := r2.URL.Query()
	r2.URL.RawQuery = q.Encode()

	return &r2
}
