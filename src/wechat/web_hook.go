package wechat

import (
	"fmt"
	"net/http"
	"errors"
	"sort"
	"crypto/sha1"
	"io"
	"strings"

	"encoding/xml"
)

type RequestData struct{
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType CdataString `xml:"MsgType"`
	Content CdataString `xml:"Content"`
	MsgId string `xml:"MsgId"`
}


type CdataString struct {
	Value string `xml:",cdata"`
}


type ReplyData struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType CdataString `xml:"MsgType"`
	Content CdataString `xml:"Content"`
}



func makeSignature(timestamp string, nonce string) string {
	sl := []string{TOKEN, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func ValidateRequest(r *http.Request) error {
	if (r.Form == nil) {
		return errors.New("Form is empty")
	}
	signature := r.Form.Get("signature")
	timestamp := r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	hashcode := makeSignature(timestamp, nonce)
	if hashcode != signature {
		return errors.New("hash value not match")
	}
	return nil
}