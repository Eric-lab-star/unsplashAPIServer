package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Eirc-lab-star/apiServer/query"
)

// "/photos" --> get all images
// "/photos?id=1" --> get iamge by id
func photos(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := r.FormValue("id")

		if id == "" {
			withoutId(w, db)
			return
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			w.Write([]byte("Invalid id"))
		}

		withId(w, db, intId)
		return
	}
}

func withoutId(w http.ResponseWriter, db *sql.DB) {
	images, err := query.GetAllImages(db)
	if err != nil {
		w.Write([]byte("No images found"))
		return
	}

	b, err := json.Marshal(images)
	if err != nil {
		fmt.Println("error marshalling json")
		panic(err)
	}
	w.Write(b)
}

func withId(w http.ResponseWriter, db *sql.DB, id int) {

	image, err := query.GetImageById(db, id)
	if err != nil {
		w.Write([]byte("Image not found"))
		return
	}

	b, err := json.Marshal(*image)
	if err != nil {
		fmt.Println("error marshalling json", err)
	}

	w.Write(b)
}
