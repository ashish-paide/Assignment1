package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
)

func getAllBlocks() (string){
	filePath := "ledger.txt"
	
	//opening file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	line := "["
	defer file.Close()

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)

	// Read line by line
	if(scanner.Scan()){
		line += scanner.Text()
	}

	for scanner.Scan() {
		line += ","
		line +=  scanner.Text()
		
	}
	line += "]"
	//fmt.Println(line)

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return line
}

func getBlockById(id int)(string){
	filePath := "ledger.txt"
	
	//opening file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		id -= 1
		fmt.Println(id)
		line := scanner.Text()
		if id == 0{
			return line
		}
		
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return "not found data"
}

	