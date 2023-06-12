package main

import (
	"time"
	"fmt"
	//"strings"
	"strconv"
	"github.com/syndtr/goleveldb/leveldb"
	"reflect"
)

func intToString(num int) string {
	return  strconv.Itoa(num)
}

type block struct{
	blockNo int32
	prevBlockHash int64
	tracsactions[] struct{
		ID int32
		version float64
		valid bool	
	}
	timeStamp time.Time
}


//defining the structure of the data in leveldb
type leveldbVal struct {
	val int64
	ver float64
}

//declaring the database and err globally
var db *leveldb.DB
var err error

func main() {
	//constructing the database
	db, err = leveldb.OpenFile("db", nil)
	if err != nil {
		fmt.Println("issue in Database OPENFILE", err)
	}

	//create the keys SIM1 --> SIM1000
	for i := 1 ; i <= 1000 ; i++{
		key := "SIM" + intToString(i) 
		err = db.Put([]byte(key), []byte("") , nil)
		if(err != nil){
			fmt.Println("whach out" , i)
		}
	}

	//checking all the keys
	// iter := db.NewIterator(nil, nil)
	// for iter.Next() {
	// 	fmt.Println(string(iter.Key()))
	// }
	// fmt.Println(reflect.TypeOf(db))

	

}

