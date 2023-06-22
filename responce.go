//this go file is to fetch the data from the ledger.txt file 


package main

import (
	"log"
	"bufio"
	"os"
	//"fmt"
)

//	Made this function take all the blocks in the ledger.txt file and 
//make it as a json marshalled string and returns the string
func getAllBlocks() (string){
	filePath := "ledger.txt"
	
	//opening file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
//string initialisation
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


//function to frtch the block with the block id
//parameters
// -id(int) id for fetching the particular block
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

	// Read line by line ignores all other lines
	for scanner.Scan() {
		id -= 1
		//fmt.Println(id)
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

	