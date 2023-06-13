package main

import(
	"github.com/syndtr/goleveldb/leveldb"
	"encoding/json"
	"encoding/csv"
	"os"
	"log"
)

type golevelInterface interface {
	NewDatabase(path string)(*golevelDatabase ,error)
	Set(key string , value leveldbVal) error                 //insert the key(string) <--> value([]byte) pair into the database 
	Get(key []byte)(*golevelDatabase , error)				 //fetch the value with  the key from the database
	GetallInCsv()(error)									 //using for the debugging :) creates the csv file contains all the key value pairs in the database
}


//struct defining the database
type golevelDatabase struct {
	db *leveldb.DB
}


//creates the database
//Parameters:
//	 -path(string) at what place we want to locate / previously located.
func Create_Database(path string)(*golevelDatabase ,  error) {
	db, err := leveldb.OpenFile(path, nil)
	return &golevelDatabase{db:db} , err
}


//Get the value from the database
//Parameters:
//	-key(string) 
func (b *golevelDatabase) Get(key string) (string , error) {
	fetched_string , err:= b.db.Get([]byte(key) , nil)
	return string(fetched_string), err
}


//insert or update the data with the key
// Parameters
// 	-key(string) with which key we want to insert into the database
// 	-value(leveldbVal (struct)) with which value we want to insert into the database
func (b *golevelDatabase) Set(key string , value leveldbVal)error{

	jsonStr , err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	return b.db.Put([]byte(key) , []byte(jsonStr) , nil)
}


//using for debugging the database
func (b *golevelDatabase)GetallInCsv()error{

	outputFile , err := os.Create("Output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"key", "value"}
	writer.Write(header)

	iter := b.db.NewIterator(nil , nil)

	for iter.Next() {
			record := []string{
				string(iter.Key()),
				string(iter.Value()),
			}
			writer.Write(record)
	}
	iter.Release()
	return iter.Error()
}