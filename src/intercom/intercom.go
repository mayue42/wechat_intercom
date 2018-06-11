package intercom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"errors"
	"io/ioutil"
)

func makeSignature(message []byte, secret string) string {
    key := []byte(secret)
    h := hmac.New(sha256.New, key)
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
	if(len(sign)!=44){
		return errors.New("X-Hub-Signature not accecpt ")
	}
	if hashcode != sign[5:] {
		return errors.New("hash value not match")
	}
	return nil
}