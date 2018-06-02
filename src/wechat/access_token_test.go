package wechat

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	token:=GetAccessToken()
    if token=="" {        //测试函数
        t.Error("error:token is null") // 如果不是如预期的那么就报错
    } else {
        t.Log("pass") //记录一些你期望记录的信息
    }
}
