package main

import (
	"time"
	"fmt"
	//"strings"
	"strconv"
	"github.com/syndtr/goleveldb/leveldb"
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



func main() {
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		fmt.Println("issue in Database OPENFILE", err)
	}

	for i := 1 ; i <= 1000 ; i++{
		key := "SIM" + intToString(i) 
		err = db.Put([]byte(key), []byte("") , nil)
		if(err != nil){
			fmt.Println("whach out" , i)
		}
	}
	fmt.Println("done with here")

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Println(string(iter.Key()))
	}

}

