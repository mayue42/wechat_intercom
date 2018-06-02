package wechat

import (
	"time"
	"sync"
)

type AccessToken struct {
	token  string
	expire int64
	sync.RWMutex
}

var access_token AccessToken
const refresh_limit=600 //refresh 10 min before expire

func GetAccessToken()string{
	now:=time.Now().Unix()
	if(now-access_token.expire<refresh_limit){
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
	if(now-access_token.expire<refresh_limit){
		//todo refresh token
	}
}