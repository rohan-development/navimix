package deemix

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
)

//var Client *http.Client

func init() {

}

func Login(arl, address string) *http.Client {
	var Client *http.Client
	jar, _ := cookiejar.New(nil)
	Client = &http.Client{Jar: jar}
	var request Deemix
	request.Arl = arl
	request.Child = 0
	request.Force = true

	// jar, _ := cookiejar.New(nil)
	// Client = &http.Client{Jar: jar}
	jsonBytes, err := json.Marshal(request)
	check_err(err)
	resp, err := Client.Post(address+"api/loginArl", "application/json", bytes.NewBuffer(jsonBytes))
	check_err(err)
	//io.ReadAll(resp.Body)
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// check_err(err)

	// // Print it
	// fmt.Println(string(body))
	return Client
}

func AddToQueue(id, address string, Client *http.Client) {
	var requests Deemix
	requests.Url = "https://deezer.com/track/" + id

	jsonBytes, err := json.Marshal(requests)
	check_err(err)
	resp, err := Client.Post(address+"api/addToQueue", "application/json", bytes.NewBuffer(jsonBytes))
	check_err(err)
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// check_err(err)

	// // Print it
	// fmt.Println(string(body))
}

func IsDone(id, address string, Client *http.Client) bool {
	resp, err := Client.Get(address + "api/getQueue")
	check_err(err)
	defer resp.Body.Close()
	var data Root
	err = json.NewDecoder(resp.Body).Decode(&data)
	if data.Queue["track_"+id+"_1"].Status == "completed" {
		return true
	}
	return false
}
