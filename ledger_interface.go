package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type TrnxController struct{
	BlockNo int
	PrevBlockHash string
	TrnxNo int
	TrnxPair chan LedgerPair
	MaxTrnx int
	//Block Block
	Blockchan chan Block
	PrintBlockChan chan Block
}

type Block struct{
	BlockNo int
	prevBlockHash string
	Transactions []map[string]LedgerTransactionData
	timeStamp time.Time
}

func (ctrl TrnxController)initialize(){
	ctrl.BlockNo = 1
	ctrl.PrevBlockHash = ""
	ctrl.TrnxNo = 1
	ctrl.TrnxPair = make(chan LedgerPair)
	ctrl.MaxTrnx = 5

	block  := Block{
		BlockNo: 1,
		prevBlockHash: "",
		Transactions: make([]map[string]LedgerTransactionData, 0),
		timeStamp: time.Now(),
	}
	ctrl.Blockchan <- block
}

func (ctrl TrnxController)BlockInsertService() {
	
	for {
	    
		trnxPair := <-ctrl.TrnxPair
		block := <- ctrl.Blockchan
		if(len(block.Transactions) == ctrl.MaxTrnx){ 
			blockByteStream , err := json.Marshal(block)
			if err != nil {
				log.Fatal(err , "error while marshalling block")
			}

			ctrl.PrevBlockHash = calculateHash(string(blockByteStream))
			block.timeStamp = time.Now()

			ctrl.PrintBlockChan <- block


			ctrl.BlockNo ++ 
			block.BlockNo = ctrl.BlockNo
			block.prevBlockHash = ctrl.PrevBlockHash
			block.Transactions = make([]map[string]LedgerTransactionData, 0)
			block.timeStamp = time.Now()		
		}

		if(len(block.Transactions) <= ctrl.MaxTrnx){ 
			block.Transactions = append(block.Transactions , map[string]LedgerTransactionData {trnxPair.Key: trnxPair.Trnx})
		}

		ctrl.Blockchan <- block





	}
}

func (ctrl TrnxController)writeFile(){
	
	for{

		block :=  <- ctrl.PrintBlockChan
		byteStream , err := json.Marshal(block)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("ledger.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
    		panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(string(byteStream)); err != nil {
    		panic(err)	
		}
	}
	


}



