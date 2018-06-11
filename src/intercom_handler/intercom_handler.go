package intercom_handler

import (
	"net/http"
	"fmt"
	"util"
	"wechat"
	"gopkg.in/intercom/intercom-go.v2"
	myintercom "intercom"
)

func HandleIntercomGet(w http.ResponseWriter, r *http.Request){

}


func HandleIntercomPost(w http.ResponseWriter, r *http.Request){
	// to do verify token
	if err:=myintercom.ValidateRequest(r);err!=nil{
		fmt.Println(err)
		return
	}


	// get content
	notif, err := intercom.NewNotification(r.Body)
	if(err!=nil){
		fmt.Println("Read body error")
		fmt.Println(err)
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