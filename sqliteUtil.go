package main

import (
	"database/sql"
	"log"

	"github.com/Eirc-lab-star/apiServer/unsplash"
)

// example of image
// data := unsplash.Image{
// 	Id:          "iksvC-BOTFg",
// 	Description: "A multi - colored building with a bridge going across it",
// 	Urls: unsplash.ImageUrl{
// 		Full:    "https://images.unsplash.com/photo-1719861032503-225fac307c59?crop=entropy&cs=srgb&fm=jpg&ixid=M3w2MzA2Njl8MHwxfGFsbHwxfHx8fHx8Mnx8MTcyMDM3NDEyOHw&ixlib=rb-4.0.3&q=85",
// 		Regular: "https://images.unsplash.com/photo-1719861032503-225fac307c59?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w2MzA2Njl8MHwxfGFsbHwxfHx8fHx8Mnx8MTcyMDM3NDEyOHw&ixlib=rb-4.0.3&q=80&w=1080",
// 		Small:   "https://images.unsplash.com/photo-1719861032503-225fac307c59?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w2MzA2Njl8MHwxfGFsbHwxfHx8fHx8Mnx8MTcyMDM3NDEyOHw&ixlib=rb-4.0.3&q=80&w=400",
// 		Thumb:   "https://images.unsplash.com/photo-1719861032503-225fac307c59?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3w2MzA2Njl8MHwxfGFsbHwxfHx8fHx8Mnx8MTcyMDM3NDEyOHw&ixlib=rb-4.0.3&q=80&w=200",
// 	},
// 	User: unsplash.User{FirstName: "Alexey", LastName: "Komissarov"},
// }

func openDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

var initImageTable = `create table if not exists images (id integer not null primary key autoincrement, fname text, lname text, desc text, urlFull text, urlReg text, urlSmall text, urlThumb text);`

func createTable(db *sql.DB) {
	_, err := db.Exec(initImageTable)
	if err != nil {
		log.Printf("%q: %s\n", err, initImageTable)
		log.Fatal(err)
	}
}

var insertImage = `insert into images (fname, lname, desc, urlFull, urlReg, urlSmall, urlThumb) values (?, ?, ?, ?, ?, ?, ?);`

func insert(db *sql.DB, data *unsplash.Image) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(insertImage)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.User.FirstName, data.User.LastName, data.Description, data.Urls.Full, data.Urls.Regular, data.Urls.Small, data.Urls.Thumb)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func printImages(db *sql.DB) {
	rows, err := db.Query("select * from images")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var fname, lname, desc, urlFull, urlReg, urlSmall, urlThumb string
		err = rows.Scan(&id, &fname, &lname, &desc, &urlFull, &urlReg, &urlSmall, &urlThumb)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, fname, lname, desc)
	}
}
