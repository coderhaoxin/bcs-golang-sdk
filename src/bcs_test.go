package bcs

import (
	"fmt"
	"os"
	"testing"
)

const (
	AK = "XXXX" //replace it with you ak
	SK = "XXXX" //replace it with you sk
)

func TestBCS(t *testing.T) {
	bcs := New(AK, SK)
	err := bcs.CreateBucket("gobucket")
	if err != nil {
		t.Errorf("create bucket error: %s", err.Error())
	}
	bl, err := bcs.ListBucket()
	if err != nil {
		t.Errorf("list bucket error: %s", err.Error())
	} else {
		for _, bi := range *bl {
			fmt.Printf("bucket name: %s\n", bi.Bucket_name)
		}
	}
	err = bcs.DeleteBucket("gobucket")
	if err != nil {
		t.Errorf("delete bucket error: %s", err.Error())
	}
	out, err := os.Create("d.txt")
	if err != nil {
		t.Errorf("open file error\n")
	} else {
		defer out.Close()
		err = bcs.GetObject("goodluck", "/t.html", out)
		if err != nil {
			t.Errorf("get object error: %s", err.Error())
		}
	}
	in, err := os.Open("d.txt")
	if err != nil {
		t.Errorf("open file error\n")
	}
	stat, err := in.Stat()
	if err != nil {
		t.Errorf("stat file error\n")
	}
	len := stat.Size()
	err = bcs.PutObject("goodluck", "/s.html", len, in)
	if err != nil {
		t.Errorf("put object error: %s", err.Error())
	}
	err = bcs.DeleteObject("goodluck", "/s.html")
	if err != nil {
		t.Errorf("delete object error: %s", err.Error())
	}
}
