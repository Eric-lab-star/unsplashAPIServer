package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Eirc-lab-star/apiServer/unsplash"
	_ "github.com/mattn/go-sqlite3"
)

var baseURI = "https://api.unsplash.com"

func main() {
	keys := parseKeys()

	db, err := sql.Open("sqlite3", "./unsplash.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table Images(id integer not null primary key, fname text, lname text, desc text, urlFull text, urlReg text urlSmall text, urlThumb text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Unsplash API"))
	})

	http.HandleFunc("/photos", func(w http.ResponseWriter, r *http.Request) {
		photo := fmt.Sprintf("%s/photos", baseURI)
		b := unsplash.Get(r, keys.Akey, photo)
		w.Write(b)

	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

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
