package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../models"
	_ "github.com/mattn/go-sqlite3"
	"goji.io/pat"
)

//Global SQL Variable
var db *sql.DB

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
	var err error
	db, err = sql.Open("sqlite3", "../jrdd.db")
	checkErr(err)

	//Fail if can't connect to DB
	checkErr(db.Ping())
}

func GetData(w http.ResponseWriter, r *http.Request) {
	search := pat.Param(r, "search")
	// Query data
	rows, err := db.Query("SELECT songs.artist, songs.song, songs.length, genres.name as genre FROM songs join genres on songs.genre = genres.id where song like? OR artist like? OR genres.name like?", search, search, search)
	checkErr(err)
	defer rows.Close()

	res := make([]models.Songs, 0)

	//Iterate through result set
	for rows.Next() {
		var artist string
		var song string
		var genre string
		var length int
		err := rows.Scan(&artist, &song, &length, &genre)
		checkErr(err)

		//Append result set to res variable
		res = append(res, models.Songs{
			Song:   song,
			Artist: artist,
			Genre:  genre,
			Length: length,
		})
	}

	jsonOut, _ := json.Marshal(res)

	//Send API response on JSON format
	fmt.Fprintf(w, string(jsonOut))

}

func GetByLength(w http.ResponseWriter, r *http.Request) {
	totalLen := pat.Param(r, "len")
	s := strings.Split(totalLen, "-")
	firstLen, secondLen := s[0], s[1]

	// Query data
	rows, err := db.Query("SELECT songs.artist, songs.song, songs.length, genres.name as genre FROM songs join genres on songs.genre = genres.id where songs.length between ? and ? order by songs.length desc", firstLen, secondLen)
	checkErr(err)
	defer rows.Close()

	res := make([]models.Songs, 0)

	//Iterate through result set
	for rows.Next() {
		var artist string
		var song string
		var genre string
		var length int
		err := rows.Scan(&artist, &song, &length, &genre)
		checkErr(err)

		//Append result set to res variable
		res = append(res, models.Songs{
			Song:   song,
			Artist: artist,
			Genre:  genre,
			Length: length,
		})
	}

	jsonOut, _ := json.Marshal(res)

	//Send API response on JSON format
	fmt.Fprintf(w, string(jsonOut))

}
