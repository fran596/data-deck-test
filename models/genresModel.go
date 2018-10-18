package models

//Model of a Genre
type Genres struct {
	Genre  string "json:genre"
	Length int    "json:length"
	Songs  int    "json:songs"
}
