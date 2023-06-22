package main

import (
	"testing"
)

func TestCreate_Database(t *testing.T) {
	db := Create_Database("db")
	
	if db == nil {
		t.Fail()
	}

	type LocalTransactionData struct {
		Value int
		Version float64
	}

	type testCase  struct {
		Key string 
		Value LocalTransactionData
	}

	

	testcases := []testCase{
		{Key : "SIM1" , Value: LocalTransactionData{Value : 1, Version : 1}} ,
		{Key : "SIM2" , Value: LocalTransactionData{Value : 2, Version : 1}} ,
		{Key : "SIM30" , Value: LocalTransactionData{Value : 1, Version : 1}} ,
		{Key : "SIM40" , Value: LocalTransactionData{Value : 1, Version : 1}} ,
		{Key : "SIM1000" , Value: LocalTransactionData{Value : 1, Version : 1}} ,
	}

	if testcases == nil {
		t.Fail()
	}

}

func TestFirstInstace(t *testing.T){
	// testing the first instance of the level db
}

func TestSetAndGetIntoDataBase(t *testing.T){
	//Testing the put and get operations in database
}



