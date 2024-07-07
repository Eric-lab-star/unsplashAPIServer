package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseURI = "https://api.unsplash.com"

type resp struct {
	Id   string `json:"name"`
	Akey string `json:"ACCESS_KEY"`
	Skey string `json:"SECRET_KEY"`
}

func main() {
	file, err := os.Open("key.json")
	info, err := file.Stat()
	s := info.Size()
	if err != nil {
		panic(err)
	}
	data := make([]byte, s)
	_, err = file.Read(data)
	env := resp{}
	err = json.Unmarshal(data, &env)
	if err != nil {
		fmt.Printf("json panic %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		photo := fmt.Sprintf("%s/photos", baseURI)
		c := &http.Client{}
		req, err := http.NewRequest("GET", photo, nil)
		req.Header.Add("Authorization", "Client-ID "+env.Akey)
		if err != nil {
			panic(err)
		}
		resp, err := c.Do(req)
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
