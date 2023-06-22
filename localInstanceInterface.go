package main

import(
	"strconv"
	"log"
)

//defining the structure of the data in leveldb
type Payload struct {
	SIM map[string]TransactionData `json:"SIM"`
}

type TransactionData struct {
	Val int     `json:"val"`
	Ver float64 `json:"ver"`
}

func (b *golevelDatabase)createKeys(){
	trnxData := TransactionData{Val:1 , Ver:1.0}
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

func (b *golevelDatabase)updateLocalDb(payload []map[string]TransactionData)(){
	for _, entry := range payload {
		for key, value := range entry {
			prevTrnxVal , err := b.Get(key)
			if err != nil{
				log.Fatal(err)
			}
			if(prevTrnxVal.Ver == value.Ver) {
			value.Ver++
			b.Set(key , value)
			}
		}
	}
	b.GetallInCsv()
}

