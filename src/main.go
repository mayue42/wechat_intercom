package main


import (
	"fmt"
	"net/http"
	"log"

	"io/ioutil"
	"encoding/xml"
	"os"
	"time"
	"wechat"
	"gopkg.in/intercom/intercom-go.v2"
	myintercom "intercom"
	"util"
)



type User struct{
	wechat_user *wechat.UserInfoReply
	intercom_user *intercom.User
}

var user_map =map[string]*User{}

var ic = intercom.NewClient(myintercom.ACCESS_TOKEN,"")



func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "hello")
}


func HandleWXGet(w http.ResponseWriter, r *http.Request){
	if err:=wechat.ValidateRequest(r); err!=nil{
		fmt.Fprint(w, err.Error())
		return
	}

	echostr := r.Form.Get("echostr")
	fmt.Fprint(w,echostr)
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

	openid:=request.FromUserName
	if(user_map[openid]==nil){
		user,err:=wechat.GetUserInfo(openid)
		if(err!=nil){
			fmt.Println(err.Error())
			return
		}
		itercom_user := intercom.User{
			UserID: openid,
			Name: user.NickName,
			UpdatedAt: int64(time.Now().Unix()),

			//CustomAttributes: map[string]interface{}{"is_cool": true},
		}
		savedUser, err := ic.Users.Save(&itercom_user)
		if(err!=nil){
			fmt.Println(err.Error())
		}
		fmt.Println(savedUser)
		user_map[openid]=&User{user,&savedUser}

		//msg := intercom.NewUserMessage(intercom.User{ID: user_map[openid].intercom_id}, request.Content.Value)
		msg := intercom.NewUserMessage(intercom.User{UserID: openid}, request.Content.Value)
		savedMessage, err := ic.Messages.Save(&msg)
		if(err!=nil){
			fmt.Println(err.Error())
			return
		}
		fmt.Println(savedMessage)
		fmt.Print("message send to intercom suc")
	}else{
		c,err:=ic.Conversations.Reply("last", user_map[openid].intercom_user,intercom.CONVERSATION_COMMENT,request.Content.Value)
		if(err!=nil){
			fmt.Println(err.Error())
			return
		}
		fmt.Println(c)
	}
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

func HandleIntercomGet(w http.ResponseWriter, r *http.Request){

}


func HandleIntercomPost(w http.ResponseWriter, r *http.Request){
	// to do verify token

	// get content
	notif, err := intercom.NewNotification(r.Body)
	if(err!=nil){
		fmt.Println(err.Error())
		return
	}
	fmt.Println(notif)
	fmt.Println(notif.Conversation)

	if(notif.Topic=="conversation.admin.replied") {
		openid := notif.Conversation.User.UserID
		mss := notif.Conversation.ConversationParts.Parts
		text := ""
		fmt.Println("message recieve:")
		for _, ms := range mss {
			fmt.Println(ms.Body)
			text += (util.RemoveTag(ms.Body) + "\n")
		}
		err = wechat.SendTextMessage(openid, text)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}



func HandleIntercom(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r)
	fmt.Println("url:")
	fmt.Println(r.URL)
	fmt.Println("body:")
	fmt.Println(r.Body)
	fmt.Println("postform:")
	fmt.Println(r.PostForm)
	if(r.Method=="GET"){
		HandleIntercomGet(w,r)
	}else if(r.Method=="POST"){
		HandleIntercomPost(w,r)
	}
}


func intercomtest(){
	defer os.Exit(0)
	fmt.Println(ic.AppID)
	fmt.Println(ic.APIKey)
	user := intercom.User{
		UserID: "27",
		Email: "test@example.com",
		Name: "InterGopher",
		SignedUpAt: int64(time.Now().Unix()),
		CustomAttributes: map[string]interface{}{"is_cool": true},
	}
	savedUser, err := ic.Users.Save(&user)
	if(err!=nil){
		fmt.Println(err.Error())
	}
	fmt.Println(savedUser)


	//payload = struct{
	//
	//	'type': 'user',
	//	'message_type': 'comment',
	//	'user_id': user_id,
	//	'body': body
	//}
	//resp = session.post(api('conversations/last/reply'), json=payload)


	//msg := intercom.NewUserMessage(intercom.User{ID: savedUser.ID}, "body123")
	msg := intercom.NewUserMessage(intercom.User{UserID: "27"}, "body123")
	savedMessage, err := ic.Messages.Save(&msg)

	if(err!=nil){
		fmt.Println(err.Error())
	}
	fmt.Println(savedMessage)

	c,err:=ic.Conversations.Reply("last", &savedUser,intercom.CONVERSATION_COMMENT,"append message")
	if(err!=nil){
		fmt.Println(err.Error())
		return
	}
	fmt.Println(c)

	l,err:=ic.Conversations.ListAll(intercom.PageParams{})
		if(err!=nil){
		fmt.Println(err.Error())
	}
	fmt.Println(l)
}


func main() {
	//intercomtest()

	http.HandleFunc("/", Index)
	http.HandleFunc("/wx", HandleWX)
	http.HandleFunc("/intercom",HandleIntercom)

	fmt.Println("listen on port 80")
	err := http.ListenAndServe(":80", nil);
	if  err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}