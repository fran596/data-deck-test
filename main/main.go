package main

import (
	"../controllers"
	"../routes"
)

func main() {
	controllers.StartDB()
	if !controllers.GetError() {
		routes.SetUpRoutes()
	}
}
