package wechat_handler

import (
	"fmt"
	"net/http"
	"wechat"
	"io/ioutil"
	"encoding/xml"
	"time"
	"gopkg.in/intercom/intercom-go.v2"
	myintercom "intercom"
)

type User struct{
	wechat_user *wechat.UserInfoReply
	intercom_user *intercom.User
}

var user_map =map[string]*User{}

func HandleWXGet(w http.ResponseWriter, r *http.Request){
	if err:=wechat.ValidateRequest(r); err!=nil{
		fmt.Fprint(w, err.Error())
		return
	}

	echostr := r.Form.Get("echostr")
	fmt.Fprint(w,echostr)
}

func createIntercomUser(openid string)error{
	//get user info from wechat
	wechat_user, err := wechat.GetUserInfo(openid)
	if (err != nil) {
		return err
	}

	//convert data struct
	intercom_user := intercom.User{
		UserID: openid,
		Name: wechat_user.NickName,
		UpdatedAt: int64(time.Now().Unix()),
	}

	//create user in intercom
	ic := intercom.NewClient(myintercom.ACCESS_TOKEN,"")
	savedUser, err := ic.Users.Save(&intercom_user)
	if (err != nil) {
		return err
	}

	//update user_map
	user_map[openid] = &User{wechat_user, &savedUser}
	return nil
}


func HandleWXPost(w http.ResponseWriter, r *http.Request){
	//check
	if err:=wechat.ValidateRequest(r); err!=nil{
		fmt.Fprint(w, err.Error())
		return
	}

	//get content
	result, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Printf("%s\n", result)
	request:=wechat.RequestData{}
	xml.Unmarshal([]byte(result),&request)
	fmt.Println(request)

	if(wechat.AUTO_REPLY) {
		reply := wechat.ReplyData{}
		reply.Content = wechat.CdataString{"消息已收到，请耐心等待回复"}
		reply.CreateTime = (time.Now().Unix())
		reply.FromUserName = request.ToUserName
		reply.ToUserName = request.FromUserName
		reply.MsgType = wechat.CdataString{"text"}
		str, err := xml.Marshal(reply)
		if (err != nil) {
			fmt.Println("server data error")
			return
		}
		w.Write(str)
	}

	ic := intercom.NewClient(myintercom.ACCESS_TOKEN,"")
	openid:=request.FromUserName.Value

	if(user_map[openid]==nil){
		item_user,err:=ic.Users.FindByUserID(openid)
		if(err==nil) {
			user_map[openid]=&User{nil,&item_user}
		}else{
			//create new user
			if err:=createIntercomUser(openid); err!=nil{
				fmt.Printf("Error: new intercom user create: %s",err)
				return
			}
			//send message
			msg := intercom.NewUserMessage(intercom.User{UserID: openid}, request.Content.Value)
			savedMessage, err := ic.Messages.Save(&msg)
			if (err != nil) {
				fmt.Printf("Error: create msg to intercom: %s", err)
				return
			}
			fmt.Printf("message send to intercom suc: %s",savedMessage)
			return
		}
	}

	//conversition already exist
	c,err:=ic.Conversations.Reply("last", intercom.User{UserID: openid}, intercom.CONVERSATION_COMMENT,request.Content.Value)
	if(err!=nil){
		fmt.Printf("Error: reply conversation to intercom: %s", err)
		return
	}
	fmt.Printf("message send to intercom suc: %s",c)
}

func HandleWX(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r)
	fmt.Println("url:")
	fmt.Println(r.URL)
	fmt.Println("body:")
	fmt.Println(r.Body)
	fmt.Println("postform:")
	fmt.Println(r.PostForm)
	if(r.Method=="GET"){
		HandleWXGet(w,r)
	}else if(r.Method=="POST"){
		HandleWXPost(w,r)
	}
}

