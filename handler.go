package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

// Handler for the POST Transactions endpoint /create
func insertTransactions(c *gin.Context) {

	// Decode the JSON payload from the request body
	var payload []map[string]TransactionData
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, entry := range payload {
		for key, value := range entry {
			fmt.Printf("Key: %s, Value: %+v\n", key, value)
		}
	}




	// Send a success response
	c.JSON(http.StatusOK, gin.H{"status": "added successfully", "message": "Data inserted successfully"})
}