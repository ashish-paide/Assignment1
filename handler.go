package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"log"
)

type Payload struct {
	SIM map[string]LocalTransactionData `json:"SIM"`
}

// Handler for the POST Transactions endpoint /create
func insertHandler(c *gin.Context) {

	// Decode the JSON payload from the request body
	var payload []map[string]LocalTransactionData
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.insertTrnx(payload)

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"status": "inserted successfully", "message": "Data inserted successfully"})
}

func resetDBHandler(c *gin.Context){
	db.createKeys()
	c.JSON(http.StatusOK, gin.H{"status": "NewBie", "message": "reset is done"})
}

func localdbPrintHandler(c *gin.Context){
	db.GetallInCsv()
	c.JSON(http.StatusOK, gin.H{"status": "Printed", "message": "check Output.Csv file"})
}


