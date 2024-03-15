package main

import (
	db "assignment2/database"
	"assignment2/routers"
)

func main() {
	db.Init()
	routers.ServerOn().Run(":8080")
}
