package wechat

import (
	"testing"
	"net/http"
	"net/url"
	"encoding/xml"
	"fmt"
	"time"
	"os"
)




func TestRequestData(t *testing.T){
	str:="<xml><ToUserName><![CDATA[gh_1ce6b93e2b4d]]></ToUserName>"+
		"<FromUserName><![CDATA[oN4RB1qOBvrkwBi9diMYeqyXE0fc]]></FromUserName>"+
		"<CreateTime>1527834515</CreateTime>"+
		"<MsgType><![CDATA[text]]></MsgType>"+
		"<Content><![CDATA[ooooooooooooooooooooooo]]></Content>"+
		"<MsgId>6561999276080967450</MsgId>"+
		"</xml>"
	request:=RequestData{}
	xml.Unmarshal([]byte(str),&request)
	request.ToUserName==""

}

func TestReplyData(t *testing.T) {
	reply:=ReplyData{}
	reply.Content=CdataString{"test"}
	reply.CreateTime=(time.Now().Unix())
	reply.FromUserName="admin001"
	reply.ToUserName="user001"
	reply.MsgType=CdataString{"text"}
	b,err:=xml.MarshalIndent(reply,"","")
	if(err!=nil){
		fmt.Println("server data error")
		return
	}
	fmt.Println(string(b))
	os.Exit(0)
}


func TestValidateRequest(t *testing.T) {
	//http://127.0.0.1/wx?signature=c1204eef817136c24000e43cbdc5851e8bfead92&echostr=aaa&timestamp=1527827016&nonce=3943144772
	r_right:=http.Request{}
	r_right.Form=url.Values{}
	r_right.Form.Add("signature","c1204eef817136c24000e43cbdc5851e8bfead92")
	r_right.Form.Add("timestamp","1527827016")
	r_right.Form.Add("nonce","3943144772")
	err:=ValidateRequest(&r_right)
	if(err!=nil){
		t.Error("false negative")
	}else {
		t.Log("pass")
	}

	r_wrong:=http.Request{}
	r_wrong.Form=url.Values{}
	r_wrong.Form.Add("signature","c1204eef817136c24000e43cbdc5851e8bfead92")
	r_wrong.Form.Add("timestamp","1527827016")
	r_wrong.Form.Add("nonce","3943144773") //last digit from 2 to 3
	err=ValidateRequest(&r_wrong)
	if(err==nil){
		t.Error("false positive")
	}else {
		t.Log("pass")
	}

	r_empty:=http.Request{}
	err=ValidateRequest(&r_empty)
	if(err==nil){
		t.Error("false positive")
	}else {
		t.Log("pass")
	}
}
