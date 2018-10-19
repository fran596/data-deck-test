# Project Title

Jr Data Deck test for BeenVerified. 
This project builds an API for making queries to a music database.

### Prerequisites

Golang, Glide package manager and GCC

### Set Up

You must install the project dependencies using glide
```
glide install
```

### Running The Project

Running the project is pretty straightforward. Using a console, navigate to the main folder
```
cd main
```
Start the project

```
go run main.go
```

### API Usage

For searching a song, genre or artist
```
http://localhost:8000/:searchTerm

Example:
http://localhost:8000/:pop
```
For getting all the songs between a length range
```
http://localhost:8000/length/:firstLen-secondLen

Example:
http://localhost:8000/length/123-168
```

For getting all the genres information
```
http://localhost:8000/genres/get

```

## Built With

* [Golang](https://golang.org/) - The language used
* [Glide](https://github.com/Masterminds/glide) - The Go package manager
* [Go-Sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go using database/sql

## Authors

* **Francisco Arroyo** 


