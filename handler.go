package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
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

func getBlocksHandler(c *gin.Context){

	data := getAllBlocks()
	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON data")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "fetched successfully", "message": data})
}

func getBlockByIdHandler(c *gin.Context){
	id_str := c.Param("id")
	fmt.Println(id_str)
	id , _ := strconv.Atoi(id_str)
	data := getBlockById(id)

	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing JSON data")
			return
		}
	c.JSON(http.StatusOK, jsonData)
}

func resetDBHandler(c *gin.Context){
	db.createKeys()
	c.JSON(http.StatusOK, gin.H{"status": "NewBie", "message": "reset is done"})
}

func localdbPrintHandler(c *gin.Context){
	db.GetallInCsv()
	c.JSON(http.StatusOK, gin.H{"status": "Printed", "message": "check Output.Csv file"})
}


