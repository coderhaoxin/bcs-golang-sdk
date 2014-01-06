package bcs

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
)

const (
	BCS_BASE_URL = "http://bcs.duapp.com/"
)

type BCS struct {
	ak string
	sk string
}

type Bucket struct {
	Bucket_name    string
	Status         string
	Cdatetime      string
	Used_capacity  string
	Total_capacity string
	Region         string
}

func New(ak, sk string) *BCS {
	bcs := BCS{ak, sk}
	return &bcs
}

func (c *BCS) CreateBucket(bucket string) error {
	url := BCS_BASE_URL + bucket + "?sign=" + GenerateSignature(c.ak, c.sk, "PUT", bucket, "/")
	resp, err := doHttpRequest("PUT", url, map[string]string{"x-bs-acl": "private"}, nil)
	if err == nil {
		resp.Body.Close()
	}
	return err
}

func (c *BCS) ListBucket() (*[]Bucket, error) {
	var bl []Bucket
	url := BCS_BASE_URL + "?sign=" + GenerateSignature(c.ak, c.sk, "GET", "", "/")
	resp, err := doHttpRequest("GET", url, nil, nil)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("Read HTTP Response Error")
		}
		err = json.Unmarshal(body, &bl)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return &bl, nil
	}
	return nil, err
}

func (c *BCS) DeleteBucket(bucket string) error {
	url := BCS_BASE_URL + bucket + "?sign=" + GenerateSignature(c.ak, c.sk, "DELETE", bucket, "/")
	resp, err := doHttpRequest("DELETE", url, nil, nil)
	if err == nil {
		resp.Body.Close()
	}
	return err
}

func (c *BCS) GetObject(bucket, object string, out io.Writer) error {
	url := BCS_BASE_URL + bucket + object + "?sign=" + GenerateSignature(c.ak, c.sk, "GET", bucket, object)
	resp, err := doHttpRequest("GET", url, nil, nil)
	if err == nil {
		io.Copy(out, resp.Body)
		resp.Body.Close()
	}
	return err
}

func (c *BCS) HeadObject(bucket, object string) error {
	return errors.New("Not Implemented!")
}

func (c *BCS) PutObject(bucket, object string, len int64, in io.Reader) error {
	url := BCS_BASE_URL + bucket + object + "?sign=" + GenerateSignature(c.ak, c.sk, "PUT", bucket, object)
	headers := make(map[string]string)
	headers["x-bs-acl"] = "private"
	headers["Content-Length"] = strconv.FormatInt(len, 10)
	resp, err := doHttpRequest("PUT", url, headers, in)
	if err == nil {
		resp.Body.Close()
	}
	return err
}

func (c *BCS) DeleteObject(bucket, object string) error {
	url := BCS_BASE_URL + bucket + object + "?sign=" + GenerateSignature(c.ak, c.sk, "DELETE", bucket, object)
	resp, err := doHttpRequest("DELETE", url, nil, nil)
	if err == nil {
		resp.Body.Close()
	}
	return err
}

func (c *BCS) PutSuperfile(object string) error {
	return errors.New("Not Implemented!")
}

func (c *BCS) CopyObject(bucket string) error {
	return errors.New("Not Implemented!")
}
