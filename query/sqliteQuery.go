package query

import (
	"database/sql"
	"log"

	"github.com/Eirc-lab-star/apiServer/unsplash"
)

func OpenDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

var initImageTable = `
create table if not exists images (
	id integer not null primary key autoincrement,
	imageId text,
	fname text,
	lname text,
	desc text,
	urlFull text,
	urlReg text,
	urlSmall text,
	urlThumb text,
	urlRaw text
);`

func CreateTable(db *sql.DB) {
	_, err := db.Exec(initImageTable)
	if err != nil {
		log.Printf("%q: %s\n", err, initImageTable)
		log.Fatal(err)
	}
}

var insertImage = `
insert into images (
	imageId,
	fname,
	lname,
	desc,
	urlFull,
	urlReg,
	urlSmall,
	urlThumb,
	urlRaw) 
	values (?, ?, ?, ?, ?, ?, ?, ?, ?);`

func Insert(db *sql.DB, data *unsplash.Image) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(insertImage)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		data.Id,
		data.User.FirstName,
		data.User.LastName,
		data.Description,
		data.Urls.Full,
		data.Urls.Regular,
		data.Urls.Small,
		data.Urls.Thumb,
		data.Urls.Raw,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetImageById(db *sql.DB, id int) (*unsplash.Image, error) {
	var index int
	image := &unsplash.Image{}
	err := db.QueryRow("select * from images where id = ?", id).Scan(
		&index,
		&image.Id,
		&image.User.FirstName,
		&image.User.LastName,
		&image.Description,
		&image.Urls.Full,
		&image.Urls.Regular,
		&image.Urls.Small,
		&image.Urls.Thumb,
		&image.Urls.Raw,
	)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func GetAllImages(db *sql.DB) ([]*unsplash.Image, error) {
	images := []*unsplash.Image{}
	rows, err := db.Query("select * from images")

	if err != nil {
		log.Println("GetAllIMages -> Error querying images: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		image := &unsplash.Image{}
		var index int
		err = rows.Scan(
			&index,
			&image.Id,
			&image.User.FirstName,
			&image.User.LastName,
			&image.Description,
			&image.Urls.Full,
			&image.Urls.Regular,
			&image.Urls.Small,
			&image.Urls.Thumb,
			&image.Urls.Raw,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}
