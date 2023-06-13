package main

import(
	"time"
)

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