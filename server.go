package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Eirc-lab-star/apiServer/unsplash"
)

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

func runServer() {

	fmt.Println("Starting API Server")
	db := openDB("./images.db")
	defer db.Close()

	createTable(db)

	http.HandleFunc("/", home)

	http.HandleFunc("/photos", photos(db))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

var baseURI = "https://api.unsplash.com"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Unsplash API"))
}

func photos(db *sql.DB) http.HandlerFunc {

	photo := fmt.Sprintf("%s/photos", baseURI)
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("/photos")
		images := unsplash.Get(r, keys.Akey, photo)
		for _, image := range images {
			go insert(db, &image)
		}
		w.Write([]byte("save image"))
	}
}
