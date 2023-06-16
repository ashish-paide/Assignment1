package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"github.com/fatih/color"
	"fmt"
)

type TrnxController struct{
	BlockNo int
	PrevBlockHash string
	TrnxNo int
	WaitTime int
	MaxTrnx int
	//Block Block
	Blockchan chan Block
	PrintBlockChan chan Block
	TrnxPair chan LedgerPair
}

type LedgerTransactionData struct {
	ID int 			`json:"Id"`
	Value int     	`json:"val"`
	Version float64 `json:"ver"`
	Hash string 	`json:"hash"`
	Valid bool 		`json:"valid"`
}

type Block struct{
	BlockNo int `json:"blockNumber"`
	PrevBlockHash string `json:"prevBlockHash"`
	Transactions []map[string]LedgerTransactionData `json:"txns"`
	TimeStamp time.Time `json:"timeStamp"`
}

func (ctrl TrnxController)initialize(){
	blockCtrl.BlockNo = 1
	blockCtrl.PrevBlockHash = ""
	blockCtrl.TrnxNo = 1
	blockCtrl.MaxTrnx = 5
	blockCtrl.WaitTime = 5
	blockCtrl.Blockchan = make(chan Block , 1)
	blockCtrl.TrnxPair = make(chan LedgerPair ,5)
	blockCtrl.PrintBlockChan = make(chan Block , 1)
	
	block  := Block{
			BlockNo: 1,
			PrevBlockHash: "",
			Transactions: make([]map[string]LedgerTransactionData, 0),
			TimeStamp: time.Now(),
	}
	fmt.Println("fmt=---------------------------------")
	select {
	case blockCtrl.Blockchan <- block:
		color.Yellow("initial Block inserted succesfully")
	default:
		color.Red("waiting room is full")
	}
	fmt.Println(blockCtrl)
	
}

func (ctrl TrnxController)trnxInsertService() {

	color.Yellow("Insert Service is running")
	go func(){
		for {

			//if len(ctrl.Blockchan == 0)
			
			color.Green("took the block")
			trnxPair := <-blockCtrl.TrnxPair
			block := <- blockCtrl.Blockchan


			if(len(block.Transactions) == blockCtrl.MaxTrnx){ 
				blockCtrl.pushBlock(block)		
			}
	
			if(len(block.Transactions) <= blockCtrl.MaxTrnx){ 
				trnxPair.Trnx.ID = blockCtrl.TrnxNo
				blockCtrl.TrnxNo += 1
				block.Transactions = append(block.Transactions , map[string]LedgerTransactionData {trnxPair.Key: trnxPair.Trnx})
			}
	
			blockCtrl.Blockchan <- block
		}
	}()
}

func (ctrl TrnxController)pushBlock(block Block)(Block){

	blockByteStream , err := json.Marshal(block)
	if err != nil {
		log.Fatal(err , "error while marshalling block")
	}
	
	blockCtrl.PrevBlockHash = calculateHash(string(blockByteStream))
	color.Red(ctrl.PrevBlockHash)
	block.TimeStamp = time.Now()
	
	blockCtrl.PrintBlockChan <- block
		
	blockCtrl.BlockNo ++ 
	block.BlockNo = blockCtrl.BlockNo
	block.PrevBlockHash = blockCtrl.PrevBlockHash
	block.Transactions = make([]map[string]LedgerTransactionData, 0)
	block.TimeStamp = time.Now()
	return block
}

func (ctrl TrnxController)writeFile(){

	color.Yellow("write file service is running")
	go func(){
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
	
			if _, err = f.WriteString(string(byteStream) + "\n"); err != nil {
				panic(err)	
			}
		}
	}()
}

func (ctrl TrnxController)autoWrite(){
	color.Yellow("auto block writer is initialised")
	go func(){
		for{
			time.Sleep(time.Second * time.Duration(blockCtrl.WaitTime))
			color.Red("timestamp")
			if(len(blockCtrl.Blockchan) !=0 ){
				block := <- blockCtrl.Blockchan
				if(len(block.Transactions)!= 0){
					block = blockCtrl.pushBlock(block)
				} else {
					color.Red("block is empty")
				}

				blockCtrl.Blockchan <- block
			} else {
				color.Red("channel is empty")
			}
		}
	}()
}



