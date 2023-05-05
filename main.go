package main

import (
	"Go-RestApi-Books/database"
	"Go-RestApi-Books/routes"
)

func main() {
	database.StartDB()
	r := routes.Routers()
	r.Run(":8080")
}
