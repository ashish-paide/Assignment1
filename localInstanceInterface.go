package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	//"sync"
)

//defining the structure of the data in leveldb
type Payload struct {
	SIM map[string]LocalTransactionData `json:"SIM"`
}

type LedgerPair struct {
	Key string
	Trnx LedgerTransactionData 
}

type LedgerTransactionData struct {
	ID int 			`json:"Id"`
	Value int     	`json:"val"`
	Version float64 `json:"ver"`
	Hash string 	`json:"hash"`
	Valid bool 		`json:"valid"`
}

type LocalTransactionData struct {
	Value int `json:"val"`
	Version float64 `json:"ver"`
}

func (b *golevelDatabase)createKeys(){
	trnxData := LocalTransactionData{Value:1 , Version:1.0}
	for i := 1 ; i <= 1000 ; i++{
		key := "SIM" + strconv.Itoa(i) 
		err := b.Set(key, trnxData)
		if err != nil{
			log.Fatal(err)
		}
	}
	b.GetallInCsv()
	return
}

func (b *golevelDatabase)insertTrnx(payload []map[string]LocalTransactionData)(){
	for _, entry := range payload {
		for key, value := range entry {
			go blockCtrl.updateLocalDB(key , value)
		
		}
	}
}

func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (ctrl TrnxController)updateLocalDB(key string , trnx LocalTransactionData){

	
	var ledTrnx LedgerTransactionData
	ledTrnx.Value , ledTrnx.Version = trnx.Value , trnx.Version
	oldTrnx, _ := db.Get(key)
	fmt.Println("updating local database " , key , trnx.Value  , trnx.Version ,"old version" ,  oldTrnx.Version )
	if(oldTrnx.Version ==  trnx.Version){
		trnx.Version += 1
		db.Set(key , trnx)

		ledTrnx.Valid = true	
	} else{
		ledTrnx.Valid = false
	}
	str := key + strconv.Itoa(trnx.Value) + strconv.FormatFloat(trnx.Version, 'E', -1, 32) 
	ledTrnx.Hash = calculateHash(str)

	ctrl.TrnxPair <- LedgerPair{Key : key , Trnx : ledTrnx}
	fmt.Println("sent Transaction")
	//b.GetallInCsv()

}



