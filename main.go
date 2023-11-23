package main

import (
	"fmt"
	"mygram/database"
	"mygram/routers"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	database.StartDB()
	r := routers.StartApp()
	r.Run(fmt.Sprintf(":%s", port))
}
