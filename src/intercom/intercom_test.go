package intercom

import (
	"fmt"
	vendor_intercom "gopkg.in/intercom/intercom-go.v2"
	"time"
	"testing"
)

func TestIntercom(t *testing.T){
	var ic = vendor_intercom.NewClient(ACCESS_TOKEN,"")
	user := vendor_intercom.User{
		UserID: "27",
		Email: "test@example.com",
		Name: "InterGopher",
		SignedUpAt: int64(time.Now().Unix()),
		CustomAttributes: map[string]interface{}{"is_cool": true},
	}
	savedUser, err := ic.Users.Save(&user)
	if(err!=nil){
		t.Error(err)
		return
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
	msg := vendor_intercom.NewUserMessage(vendor_intercom.User{UserID: "27"}, "body123")
	savedMessage, err := ic.Messages.Save(&msg)
	if(err!=nil){
		t.Error(err)
		return
	}
	fmt.Println(savedMessage)

	c,err:=ic.Conversations.Reply("last", &savedUser,vendor_intercom.CONVERSATION_COMMENT,"append message")
	if(err!=nil){
		t.Error(err)
		return
	}
	fmt.Println(c)

	//l,err:=ic.Conversations.ListAll(intercom.PageParams{})
	//if(err!=nil){
	//	t.Error(err)
	//	return
	//}
	//fmt.Println(l)
}

func TestFindExistUser(t *testing.T)  {
	ic := vendor_intercom.NewClient(ACCESS_TOKEN,"")
	openid:="27"
	user,err:=ic.Users.FindByUserID(openid)
	if(err!=nil){
		t.Errorf("User not find: %s",err)
	}else if(user.UserID!="27"){
		t.Errorf("UserId do not match: %s",user.UserID)
	}
}
