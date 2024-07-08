package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ImageUrl struct {
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

func addAuthHeader(req *http.Request, akey string) {
	authHeader := fmt.Sprintf("Client-ID %s", akey)
	req.Header.Add("Authorization", authHeader)
}

func FilterResp(data []byte) []Image {

	var res []Image
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("Error unmarshalling data:")
		panic(err)
	}

	return res
}

func Get(req *http.Request, akey string, uri string) []Image {
	c := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	addAuthHeader(req, akey)

	resp, err := c.Do(req)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return FilterResp(b)
}
