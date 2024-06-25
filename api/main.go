package main

import (
	"log"
	"paleta-api/database"
	"paleta-api/routes"
)

func main() {
	log.Println("Starting api application...")
	database.ConnectDatabase()
	r:= routes.SetupRouter()
	r.Run(":8080")
}
