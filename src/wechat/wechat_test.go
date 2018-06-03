package wechat

import (
	"testing"
	"fmt"
)

const open_id="oN4RB1qOBvrkwBi9diMYeqyXE0fc"

func TestSendTextMessage(t *testing.T) {
	err:=SendTextMessage(open_id,"test123")
    if err==nil{
        t.Log("pass")
    } else {
        t.Error("error")
    }
}

func TestGetUserInfo(t *testing.T) {
	reply,err:=GetUserInfo(open_id)
	if(err==nil && reply!=nil){
		fmt.Println(reply)
        t.Log("pass")
	}else{
		fmt.Errorf(err.Error())
        t.Error("error")
	}
}
