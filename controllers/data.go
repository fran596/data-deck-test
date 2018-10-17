package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var dbError = false

func checkErr(err error, args ...string) {
	if err != nil {
		fmt.Println("Error")
		fmt.Println("%q: %s", err, args)
		dbError = true
	} else {
		dbError = false
	}

}

func GetError() bool {
	return dbError
}

func StartDB() {
	db, err := sql.Open("sqlite3", "../jrdd.db")
	checkErr(err)
	defer db.Close()

	//2. fail-fast if can't connect to DB
	checkErr(db.Ping())
}

// func getData() {
// 	// Query data

// 	rows, err := db.Query("SELECT songs.artist, songs.song, songs.length, genres.name as genre FROM songs join genres on songs.genre = genres.id where song = '' OR artist = '' OR genres.name ='Pop'")
// 	checkErr(err)
// 	defer rows.Close()

// 	//5.1 Iterate through result set
// 	for rows.Next() {
// 		var artist string
// 		var song string
// 		var genre string
// 		var length int
// 		err := rows.Scan(&artist, &song, &length, &genre)
// 		checkErr(err)
// 		fmt.Printf("name=%s\n", song)
// 	}
// }
