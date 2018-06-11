package intercom

import (
	"fmt"
	"gopkg.in/intercom/intercom-go.v2"
	"time"
	"testing"
)

func TestIntercom(t *testing.T){
	var ic = intercom.NewClient(ACCESS_TOKEN,"")
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
