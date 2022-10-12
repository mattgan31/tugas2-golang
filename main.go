package main

import (
	// "golang-tugas2/database"
	"golang-tugas2/database"
	"golang-tugas2/routers"
)

var PORT = ":3000"

func main() {

	database.StartDB()

	routers.StartServer().Run(PORT)
}
