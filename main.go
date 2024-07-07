package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Eirc-lab-star/apiServer/unsplash"
)

var baseURI = "https://api.unsplash.com"

type authKeys struct {
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
	env := authKeys{}
	err = json.Unmarshal(data, &env)
	if err != nil {
		fmt.Printf("json panic %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		photo := fmt.Sprintf("%s/photos", baseURI)
		b := unsplash.Get(r, env.Akey, photo)
		w.Write(b)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
