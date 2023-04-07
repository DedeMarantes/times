package main

import (
	"github.com/DedeMarantes/times/database"
	"github.com/DedeMarantes/times/routes"
)

func main() {
	database.ConnectDb()
	routes.HandleRequests()
}
