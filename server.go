package main

import (
	"log"
	"github.com/gin-gonic/gin"
)




var db = Create_Database("db")
func main() {

	// Create a new Gin router
	router := gin.Default()
	
	// POST /create/{key}
	go router.POST("/insert", insertHandler)
	go router.GET("/admin/reset" , resetDBHandler)


	log.Println("Server listening on port 18080...")
	router.Run(":18080")
}




