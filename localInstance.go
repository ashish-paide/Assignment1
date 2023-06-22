package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"github.com/fatih/color"
	//"sync"
)

//defining the structure of the data in leveldb
type LedgerPair struct {
	Key string
	Trnx LedgerTransactionData 
}


//defines the structure of the local transaction data structure
type LocalTransactionData struct {
	Value int `json:"val"`
	Version float64 `json:"ver"`
}


//function to create the first instance of the local database
func (b *golevelDatabase)createKeys(){
	trnxData := LocalTransactionData{Value:1 , Version:1.0}
	for i := 1 ; i <= 1000 ; i++{
		key := "SIM" + strconv.Itoa(i) 
		trnxData.Value = i
		err := b.Set(key, trnxData)
		if err != nil{
			log.Fatal(err)
		}
	}
	b.GetallInCsv()
	return
}


//function to insert all the transactions in the payload concurrently
func (b *golevelDatabase)insertTrnx(payload []map[string]LocalTransactionData)(){
	color.Yellow("about to insert transactions congurrently")
	for _, entry := range payload {
		for key, value := range entry {
			go b.updateLocalDB(key , value)
		
		}
	}
}

//function to create the hash of the string
//parameters:
//     -data(string) the string that we want to hash
func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

//function to check the validity of the transaction , calculate hash , making a LedgerTransaction Data to push into the block
func (b *golevelDatabase)updateLocalDB(key string , trnx LocalTransactionData){

	var ledTrnx LedgerTransactionData
	ledTrnx.Value , ledTrnx.Version = trnx.Value , trnx.Version
	oldTrnx, _ := b.Get(key)
	color.Green(key +  " took from the db")
	
	//checking the validity of the transaction
	if(oldTrnx.Version ==  trnx.Version){
		trnx.Version += 1
		b.Set(key , trnx)

		ledTrnx.Valid = true	
	} else{
		ledTrnx.Valid = false
	}

	//calculating the hash of the transaction
	str := key + strconv.Itoa(trnx.Value) + strconv.FormatFloat(trnx.Version, 'E', -1, 32) 
	ledTrnx.Hash = calculateHash(str)

	//defining the pair of key and transaction to push to the channel
	ledPair  := LedgerPair{
		Key : key,
		Trnx: ledTrnx,
	}

	//checking if there is any deadlock
	select {
	case blockCtrl.TrnxPair <- ledPair:
		color.Green(key + "transaction sent to ledger channel successfully /n" )
	default:
		color.Red(key + "waiting ledger transaction channel  is full /n")
	}

}



