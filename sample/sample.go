package main

import (
	"fmt"
	"../src/"
)

const (
	AK = "XXXXXXXXXXXXXXXXXX"	//replace it with you ak
	SK = "XXXXXXXXXXXXXXXXXX"	//replace it with you ak
)

func main() {
	bcs := bcs.New(AK, SK)
	err := bcs.CreateBucket("raretestbucket")
	if err != nil {
		fmt.Printf("create bucket error: %s", err.Error())
	}
	err = bcs.DeleteBucket("raretestbucket")
	if err != nil {
		fmt.Printf("delete bucket error: %s", err.Error())
	}
}
