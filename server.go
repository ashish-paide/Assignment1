package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	SIM map[string]TransactionData `json:"SIM"`
}

type TransactionData struct {
	Val int     `json:"val"`
	Ver float64 `json:"ver"`
}



func main() {

	// Create a new Gin router
	router := gin.Default()

	// POST /create/{key}
	go router.POST("/create", insertTransactions)

	log.Println("Server listening on port 18080...")
	router.Run(":18080")
}




