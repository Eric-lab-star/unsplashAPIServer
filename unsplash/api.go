package unsplash

import (
	"fmt"
	"io"
	"net/http"
)

func addAuthHeader(req *http.Request, akey string) {
	authHeader := fmt.Sprintf("Client-ID %s", akey)
	req.Header.Add("Authorization", authHeader)
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
	return b
}
