package models

//Model of a song
type Songs struct {
	Song   string "json:song"
	Artist string "json:artist"
	Genre  string "json:genre"
	Length int    "json:length"
}
