package wechat

import (
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	err:=SendTextMessage("oN4RB1qOBvrkwBi9diMYeqyXE0fc","test123")
    if err!=nil{
        t.Error("error")
    } else {
        t.Log("pass")
    }
}
