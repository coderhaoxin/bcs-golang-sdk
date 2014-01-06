package bcs

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func HashHmac(content, key string) []byte {
	keybytes := []byte(key)
	contentbytes := []byte(content)
	mac := hmac.New(sha1.New, keybytes)
	mac.Write(contentbytes)
	return mac.Sum(nil)
}

func Base64Encode(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

func UrlEncode(raw_url string) string {
	return url.QueryEscape(raw_url)
}

func GenerateSignature(ak, sk, method, bucket, object string) string {
	flag := "MBO"
	content := "Method=" + method + "\nBucket=" + bucket + "\nObject=" + object + "\n"
	return flag + ":" + ak + ":" + UrlEncode(Base64Encode(HashHmac(flag+"\n"+content, sk)))
}

func doHttpRequest(method, url string, headers map[string]string, in io.Reader) (resp *http.Response, err error) {
	req, rerr := http.NewRequest(method, url, in)
	if rerr != nil {
		return nil, errors.New("http.NewRequest Error!")
	}
	for k, v := range headers {
		if strings.ToLower(k) == "content-length" {
			req.ContentLength, _ = strconv.ParseInt(v, 10, 64)
		} else {
			req.Header.Add(k, v)
		}
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("do HTTP Request Error!")
	}
	if resp.StatusCode != 200 {
		var errmsg string
		body, rerr := ioutil.ReadAll(resp.Body)
		if rerr != nil {
			errmsg = "Read HTTP Response Error!"
		}
		errmsg = string(body)
		resp.Body.Close()
		return nil, errors.New(errmsg)
	}
	return resp, nil
}
