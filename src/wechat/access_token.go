package wechat

import (
	"time"
	"sync"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type AccessToken struct {
	token  string
	expire int64
	sync.RWMutex
}
type AccessTokenReply struct{
	Access_token string `json:"access_token"`
	Expires_in int64 `json:"expires_in"`
}

var access_token AccessToken
const refresh_limit=600 //refresh 10 min before expire

func GetAccessToken()string{
	now:=time.Now().Unix()
	if(access_token.expire-now<refresh_limit){
		refreshToken()
	}
	access_token.RLock()
	defer access_token.RUnlock()
	return access_token.token
}

func refreshToken(){
	now:=time.Now().Unix()
	access_token.Lock()
	defer access_token.Unlock()
	if(access_token.expire-now<refresh_limit){
		//todo refresh token
		//resp,err:=http.Post(TOKEN_URL,"application/x-www-form-urlencoded",nil)
		resp,err:=http.Post(TOKEN_URL,"application/json",nil)

		if(err!=nil){
			fmt.Println("error")
			fmt.Errorf(err.Error())
			return
		}
		defer resp.Body.Close()
		body,err:=ioutil.ReadAll(resp.Body)
		if(err!=nil){
			fmt.Println("error")
			fmt.Errorf(err.Error())
			return
		}
		if bytes.Contains(body, []byte("errmsg")) {
			reply:=Reply{}
			json.Unmarshal(body,&reply)
			fmt.Printf("error:%d;%s",reply.ErrCode,reply.ErrMsg)
			fmt.Errorf("%d;%s",reply.ErrCode,reply.ErrMsg)
			return
		} else {
			reply:=AccessTokenReply{}
			json.Unmarshal(body,&reply)
			access_token.token=reply.Access_token
			access_token.expire=reply.Expires_in+now
		}

	}
}