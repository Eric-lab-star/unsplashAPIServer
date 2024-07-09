package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Eirc-lab-star/apiServer/query"
	"github.com/Eirc-lab-star/apiServer/unsplash"
)

func updatePhoto(db *sql.DB) http.HandlerFunc {
	var wg sync.WaitGroup

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("/update")
		images := unsplash.GetImages(r)

		for index, image := range images {
			wg.Add(1)
			i := image
			index := index
			go func() {
				defer wg.Done()
				err := query.Insert(db, &i)
				if err != nil {
					log.Printf("image %v: ", index)
					log.Println(err)
				}
			}()
		}

		wg.Wait()
		len := len(images)
		msg := fmt.Sprintf("Saved %v images", len)
		w.Write([]byte(msg))
	}
}
