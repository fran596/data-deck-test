package routes

import (
	"fmt"
	"net/http"

	"goji.io"
)

func SetUpRoutes() {
	mux := goji.NewMux()
	// setRoutes(*mux)
	fmt.Println("Server listening on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}

// func setRoutes(mux goji.Mux) {
// 	mux.HandleFunc(pat.Get("/hello/:name"), hello)
// }
