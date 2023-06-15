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
	color.Yellow("about to insert transactions congurrently")
	for _, entry := range payload {
		for key, value := range entry {
			go b.updateLocalDB(key , value)
		
		}
	}
}

func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (b *golevelDatabase)updateLocalDB(key string , trnx LocalTransactionData){

	
	var ledTrnx LedgerTransactionData
	ledTrnx.Value , ledTrnx.Version = trnx.Value , trnx.Version
	oldTrnx, _ := b.Get(key)
	color.Yellow(key , "took from the db")
	
	fmt.Println("updating local database " , key , trnx.Value  , trnx.Version ,"old version" ,  oldTrnx.Version )
	if(oldTrnx.Version ==  trnx.Version){
		trnx.Version += 1
		b.Set(key , trnx)

		ledTrnx.Valid = true	
	} else{
		ledTrnx.Valid = false
	}
	str := key + strconv.Itoa(trnx.Value) + strconv.FormatFloat(trnx.Version, 'E', -1, 32) 
	ledTrnx.Hash = calculateHash(str)

	ledPair  := LedgerPair{
		Key : key,
		Trnx: ledTrnx,
	}
	//fmt.Println(len(blockCtrl.TrnxPair) , ctrl.MaxTrnx)
	select {
	case blockCtrl.TrnxPair <- ledPair:
		color.Yellow(key , "transaction sent to ledger channel successfully /n" )
	default:
		color.Red(key , "waiting ledger transaction channel  is full /n")
	}

}



