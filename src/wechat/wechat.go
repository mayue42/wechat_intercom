package wechat

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
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
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func SendTextMessage(open_id string, message string)error{
	token:=GetAccessToken()
	url:=fmt.Sprintf(STAFF_SEND_URL,token)
	content:=SendMessageRequest{}
	content.ToUser=open_id
	content.Text.Content=message
	content.MsgType="text"
	str,err:=json.Marshal(content)
	resp,err:=http.Post(url,"application/json",strings.NewReader(string(str)))
	if(err!=nil){
		fmt.Errorf(err.Error())
		return err
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if(err!=nil){
		fmt.Errorf(err.Error())
		return err
	}
	reply:=Reply{}
	json.Unmarshal(body,&reply)
	if(reply.ErrCode!=0 || reply.ErrMsg!="ok"){
		fmt.Errorf("reply error! code:%d;msg:%s\n",reply.ErrCode,reply.ErrMsg)
		return errors.New("reply eror")
	}
	return nil
}

func MyPrint(){
	fmt.Println("test")
}

