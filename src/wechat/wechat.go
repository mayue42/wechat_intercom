package wechat

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"bytes"
)

type Text struct{
	Content string `json:"content"`
}

type SendMessageRequest struct{
	ToUser string `json:"touser"`
	MsgType string `json:"msgtype"`
	Text Text `json:"text"`
}

type Reply struct{
	ErrCode int	 			`json:"errcode"`
	ErrMsg string 			`json:"errmsg"`
}

type UserInfoReply struct{
    Subscribe int 			`json:"subscribe"`
    Openid string 			`json:"openid"`
    NickName string 		`json:"nickname"`
    Sex int 				`json:"sex"`
    Language string 		`json:"language"`
    City string 			`json:"city"`
    Province string 		`json:"province"`
    Country string 			`json:"country"`
    Headimgurl string		`json:"headimgurl"`
    Subscribe_time int64	`json:"subscribe_time"`
    Unionid string			`json:"unionid"`
    Remark string			`json:"remark"`
    Groupid int				`json:"groupid"`
    Tagid_list []int		`json:"tagid_list"`
    Subscribe_scene string	`json:"subscribe_scene"`
    Qr_scene  int64			`json:"qr_scene"`
    Qr_scene_str string		`json:"qr_scene_str"`
}

func SendTextMessage(open_id string, message string)error{
	token:=GetAccessToken()
	url:=fmt.Sprintf(STAFF_SEND_URL,token)
	content:=SendMessageRequest{}
	content.ToUser=open_id
	content.Text.Content=message
	content.MsgType="text"
	str,err:=json.Marshal(content)
	//resp,err:=http.Post(url,"application/json",strings.NewReader(string(str)))
	//resp,err:=http.Post(url,"application/json",bytes.NewBuffer(str))
	resp,err:=http.Post(url,"application/json",bytes.NewReader(str))
	if(err!=nil){
		return err
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if(err!=nil){
		return err
	}
	reply:=Reply{}
	json.Unmarshal(body,&reply)
	if(reply.ErrCode!=0 || reply.ErrMsg!="ok"){
		return errors.New(fmt.Sprintf("reply error! code:%d;msg:%s\n",reply.ErrCode,reply.ErrMsg))
	}
	return nil
}

func GetUserInfo(open_id string)(*UserInfoReply,error){
	token:=GetAccessToken()
	url:=fmt.Sprintf(USER_INFO_URL,token,open_id)
	resp,err:=http.Post(url,"application/json",nil)
	if(err!=nil){
		return nil,err
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if(err!=nil){
		return nil,err
	}
	fmt.Println(string(body))
	if bytes.Contains(body, []byte("errmsg")) {
		reply:=Reply{}
		json.Unmarshal(body,&reply)
		err_str:=fmt.Sprintf("error:%d;%s",reply.ErrCode,reply.ErrMsg)
		fmt.Println(err_str)
		return nil,errors.New(err_str)
	} else {
		reply := UserInfoReply{}
		err = json.Unmarshal(body, &reply)
		fmt.Println(reply)
		if (err != nil) {
			return nil, err
		}
		return &reply, nil
	}
}


