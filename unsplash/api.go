package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Urls        struct {
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user"`
}

func addAuthHeader(req *http.Request, akey string) {
	authHeader := fmt.Sprintf("Client-ID %s", akey)
	req.Header.Add("Authorization", authHeader)
}

func FilterResp(data []byte) []Response {

	var res []Response
	err := json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("Error unmarshalling data:")
		panic(err)
	}

	return res
}

func Get(req *http.Request, akey string, uri string) []byte {
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

	res := FilterResp(b)

	d, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return d
}
