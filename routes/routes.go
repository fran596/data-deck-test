package routes

import (
	"fmt"
	"net/http"

	"../controllers"
	"goji.io"
	"goji.io/pat"
)

func SetUpRoutes() {
	mux := goji.NewMux()
	setRoutes(mux)
	fmt.Println("Server listening on port 8000")
	http.ListenAndServe("localhost:8000", mux)
}

func setRoutes(mux *goji.Mux) {
	mux.HandleFunc(pat.Get("/:search"), controllers.GetData)
}
