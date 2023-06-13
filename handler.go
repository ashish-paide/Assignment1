package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"log"
)

// Handler for the POST Transactions endpoint /create
func insertHandler(c *gin.Context) {

	// Decode the JSON payload from the request body
	var payload []map[string]TransactionData
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.updateLocalDb(payload)

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"status": "inserted successfully", "message": "Data inserted successfully"})
}

func resetDBHandler(c *gin.Context){
	db.createKeys()
	c.JSON(http.StatusOK, gin.H{"status": "NewBie", "message": "reset is done"})
}

