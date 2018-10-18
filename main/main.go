package main

import (
	"../controllers"
	"../routes"
)

func main() {
	//Start the DB
	controllers.StartDB()
	if !controllers.GetError() {
		//If no errors setup the API routes
		routes.SetUpRoutes()
	}
}
