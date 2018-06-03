package wechat

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	token1:=GetAccessToken()
	token2:=GetAccessToken()
    if token1=="" || token2=="" || token1!=token2{
        t.Error("error:token is null")
    } else {
        t.Log("pass")
    }
}
