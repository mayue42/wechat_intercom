package main


import (
	"fmt"
	"net/http"
	"log"

	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"io/ioutil"
	"encoding/xml"
	"os"
	"time"
	"wechat"
	intercom "gopkg.in/intercom/intercom-go.v2"
	myintercom "intercom"
	"util"
)


func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.RequestURI)
	fmt.Println(r)
	fmt.Fprintf(w, "hello")
}


func makeSignature(timestamp string, nonce string) string {
	sl := []string{wechat.TOKEN, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
type RequestData struct{
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType CdataString `xml:"MsgType"`
	Content CdataString `xml:"Content"`
	MsgId string `xml:"MsgId"`
}


type CdataString struct {
	Value string `xml:",cdata"`
}


type ReplyData struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType CdataString `xml:"MsgType"`
	Content CdataString `xml:"Content"`
}


//http://127.0.0.1/wx?signature=c1204eef817136c24000e43cbdc5851e8bfead92&echostr=aaa&timestamp=1527827016&nonce=3943144772
func HandleWXGet(w http.ResponseWriter, r *http.Request){
	if(r.Form==nil){
		fmt.Fprintf(w, "parameter need")
	}
	signature:= r.Form.Get("signature")
	timestamp:= r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echostr := r.Form.Get("echostr")

	hashcode := makeSignature(timestamp,nonce)
	if hashcode == signature{
		fmt.Fprintf(w,echostr)
	}else {
		fmt.Fprintf(w, "parameter error")
	}
}

type User struct{
	wechat_user *wechat.UserInfoReply
	intercom_user *intercom.User
}

var user_map =map[string]*User{}

var ic = intercom.NewClient(myintercom.ACCESS_TOKEN,"")

func HandleWXPost(w http.ResponseWriter, r *http.Request){
	//check
	if(r.Form==nil){
		fmt.Fprintf(w, "parameter need")
	}
	signature:= r.Form.Get("signature")
	timestamp:= r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	hashcode := makeSignature(timestamp,nonce)
	if hashcode != signature{
		fmt.Fprintf(w, "parameter error")
		return
	}

	//get content
	result, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Printf("%s\n", result)
	request:=RequestData{}
	xml.Unmarshal([]byte(result),&request)
	fmt.Println(request)

	reply:=ReplyData{}
	reply.Content=CdataString{"消息已收到，请耐心等待回复"}
	reply.CreateTime=(time.Now().Unix())
	reply.FromUserName=request.ToUserName
	reply.ToUserName=request.FromUserName
	reply.MsgType=CdataString{"text"}
	str,err:=xml.Marshal(reply)
	if(err!=nil){
		fmt.Println("server data error")
		return
	}
	w.Write(str)
	//body := bytes.NewBuffer([]byte(str))
	//fmt.Println(body)
	//to do process
	//fmt.Fprintf(w,body)
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





func xmltest(){
	str:="<xml><ToUserName><![CDATA[gh_1ce6b93e2b4d]]></ToUserName>"+
	"<FromUserName><![CDATA[oN4RB1qOBvrkwBi9diMYeqyXE0fc]]></FromUserName>"+
	"<CreateTime>1527834515</CreateTime>"+
	"<MsgType><![CDATA[text]]></MsgType>"+
	"<Content><![CDATA[ooooooooooooooooooooooo]]></Content>"+
	"<MsgId>6561999276080967450</MsgId>"+
	"</xml>"
	request:=RequestData{}
	xml.Unmarshal([]byte(str),&request)
	fmt.Println(request)

	reply:=ReplyData{}
	reply.Content=CdataString{"test"}
	reply.CreateTime=(time.Now().Unix())
	reply.FromUserName=request.ToUserName
	reply.ToUserName=request.FromUserName
	reply.MsgType=CdataString{"text"}
	b,err:=xml.MarshalIndent(reply,"","")
	if(err!=nil){
		fmt.Println("server data error")
		return
	}
	fmt.Println(string(b))
	os.Exit(0)
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