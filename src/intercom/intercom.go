package intercom

import (
	"crypto/hmac"
	"encoding/hex"
	"net/http"
	"errors"
	"io/ioutil"
	"fmt"
	"crypto/sha1"
)

func makeSignature(message []byte, secret string) string {
    key := []byte(secret)
    h := hmac.New(sha1.New, key)
    h.Write(message)
    return hex.EncodeToString(h.Sum(nil))
}


func ValidateRequest(r *http.Request) error {
	body,err:=ioutil.ReadAll(r.Body) // if read body here, than no data to read later
	if(err!=nil){
		return errors.New("Payload Reading Error")
	}
	hashcode := makeSignature(body,HUB_SECRET)

	sign:= r.Header.Get("X-Hub-Signature")
	if(sign==""){
		return errors.New("X-Hub-Signature Not find")
	}
	if(len(sign)!=45){
		return errors.New("X-Hub-Signature not accecpt ")
	}
	if hashcode != sign[5:] {
		fmt.Println(hashcode)
		fmt.Println(sign[5:])
		return errors.New("hash value not match")
	}
	return nil
}