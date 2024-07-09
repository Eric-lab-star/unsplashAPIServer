package server

import (
	"fmt"
	"net/http"

	"github.com/Eirc-lab-star/apiServer/query"
)

func Start() {
	fmt.Println("Starting API Server")
	db := query.OpenDB("./images.db")
	defer db.Close()

	query.CreateTable(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All photos are retrieved from Unsplsh api"))
	})

	http.HandleFunc("/photos", photos(db))

	http.HandleFunc("/photos/update", updatePhoto(db))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
