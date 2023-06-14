package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)




var db = Create_Database("db")
var blockCtrl TrnxController
func main() {

	

	// Create a new Gin router
	router := gin.Default()
	// POST /create/{key}
	go router.POST("/insert", insertHandler)
	go router.GET("/admin/reset" , resetDBHandler)
	log.Println("Server listening on port 18080...")
	router.Run(":18080")

	blockCtrl.initialize()
	go blockCtrl.BlockInsertService()
	go blockCtrl.writeFile()

	fmt.Println("all services are created")
}




