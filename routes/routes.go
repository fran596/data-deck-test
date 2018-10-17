package routes

import (
	"fmt"
	"net/http"

	"goji.io"
)

func SetRoutes() {
	mux := goji.NewMux()
	// mux.HandleFunc(pat.Get("/hello/:name"), hello)
	fmt.Println("Server listening on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}
