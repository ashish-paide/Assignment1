package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

var db = Create_Database("db")
var blockCtrl TrnxController

func main() {
	blockCtrl.initialize()
	fmt.Println(blockCtrl)
	blockCtrl.trnxInsertService()
	blockCtrl.writeFile()
	blockCtrl.autoWrite()
	fmt.Println(blockCtrl)
	
	// Create a new Gin router
	router := gin.Default()
	// POST /create/{key}
	go router.POST("/insert", insertHandler)
	go router.GET("/admin/reset" , resetDBHandler)
	go router.GET("/admin/getBlocks" , getBlocksHandler)
	go router.GET("/admin/getBlock/:id" , getBlockByIdHandler)
	go router.GET("/admin/printlocaldb" , localdbPrintHandler)
	log.Println("Server listening on port 18085...")
	router.Run(":18085")

	
	

	fmt.Println("all services are created")
}




