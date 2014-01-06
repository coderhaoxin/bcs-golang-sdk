package main

import (
	"../bcs"
	"fmt"
	"os"
)

const (
	AK = "your Access Key(AK)"
	SK = "your Secret Key(SK)"
)

func main() {
	bcs := bcs.New(AK, SK)

	// create bucket
	err := bcs.CreateBucket("raretestbucket")
	if err != nil {
		fmt.Printf("create bucket error: %s", err.Error())
	}

	// put object
	in, _ := os.Open("sample.go")
	stat, _ := in.Stat()
	len := stat.Size()
	err = bcs.PutObject("raretestbucket", "/sample.go", len, in)
	if err != nil {
		fmt.Printf("put object error: %s", err.Error)
	}

	// delete object
	err = bcs.DeleteObject("raretestbucket", "/sample.go")
	if err != nil {
		fmt.Printf("delete object error: %s", err.Error())
	}

	// delete bucket
	err = bcs.DeleteBucket("raretestbucket")
	if err != nil {
		fmt.Printf("delete bucket error: %s", err.Error())
	}
}
