package main

import(
	"strconv"
)

//defining the structure of the data in leveldb
type leveldbVal struct {
	Value int `json:"val"`
	Version float64  `json:"ver"`
}

func createKeys(b *golevelDatabase) error {
	for i := 1 ; i <= 1000 ; i++{
		key := "SIM" + strconv.Itoa(i) 
		err := b.db.Put([]byte(key), []byte("") , nil)
		return err
	}
	return nil
}