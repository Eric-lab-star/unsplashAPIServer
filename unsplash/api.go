package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ImageUrl struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Image struct {
	Id          string   `json:"id"`
	Description string   `json:"alt_description"`
	Urls        ImageUrl `json:"urls"`
	User        User     `json:"user"`
}

var keys = parseKeys()

type authKeys struct {
	Id   string `json:"name"`
	Akey string `json:"ACCESS_KEY"`
	Skey string `json:"SECRET_KEY"`
}

func parseKeys() authKeys {
	file, err := os.Open("key.json")
	info, err := file.Stat()
	s := info.Size()
	if err != nil {
		panic(err)
	}
	data := make([]byte, s)
	_, err = file.Read(data)
	keys := authKeys{}
	err = json.Unmarshal(data, &keys)
	if err != nil {
		fmt.Printf("json panic %v\n", err)
	}
	return keys
}

var BASEURI = "https://api.unsplash.com"

func GetImages(req *http.Request) []Image {
	uri := fmt.Sprintf("%s%s", BASEURI, "/photos?page=4&per_page=30")
	c := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}

	addAuthHeader(req, keys.Akey)

	resp, err := c.Do(req)
	remain := resp.Header.Get("X-Ratelimit-Remaining")
	log.Printf("Remained Api Request: %v", remain)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return filterRes(b)
}

func addAuthHeader(req *http.Request, akey string) {
	authHeader := fmt.Sprintf("Client-ID %s", akey)
	req.Header.Add("Authorization", authHeader)
}

func filterRes(data []byte) []Image {

	var res []Image
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("Error unmarshalling data:")
		panic(err)
	}

	return res
}
