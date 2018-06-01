package main


import (
	"fmt"
	"net/http"
	"log"

	"sort"
	"crypto/sha1"
	"io"
	"strings"
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

func HandleWX(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Body)
	fmt.Println(r.URL)
	if(r.Method=="GET"){
		HandleWXGet(w,r)
	}else if(r.Method=="POST"){
		HandleWXPost(w,r)
	}
}



func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/wx", HandleWX)

	fmt.Println("listen on port 80")
	err := http.ListenAndServe(":80", nil);
	if  err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}