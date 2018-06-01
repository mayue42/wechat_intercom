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
)

const token string="token123"


func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.RequestURI)
	fmt.Println(r)
	fmt.Fprintf(w, "hello")
}


func makeSignature(timestamp string, nonce string) string {
	sl := []string{token, timestamp, nonce}
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
	MsgType xml.CharData `xml:"MsgType"`
	Content xml.CharData `xml:"Content"`
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
	reply.Content=CdataString{"test"}
	reply.CreateTime=(time.Now().Unix())
	reply.FromUserName=request.ToUserName
	reply.ToUserName=request.FromUserName
	reply.MsgType=CdataString{"text"}
	str,err:=xml.Marshal(reply)
	if(err!=nil){
		fmt.Println("server data error")
		return
	}
	//body := bytes.NewBuffer([]byte(str))
	//fmt.Println(body)
	//to do process
	//fmt.Fprintf(w,body)
	w.Write(str)
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

func test(){
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

func main() {
	//test()
	http.HandleFunc("/", Index)
	http.HandleFunc("/wx", HandleWX)

	fmt.Println("listen on port 80")
	err := http.ListenAndServe(":80", nil);
	if  err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}