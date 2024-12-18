package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/syedazeez337/jobportalgo/internal/repository"
)


func main() {
	
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.Run(":8080")
}