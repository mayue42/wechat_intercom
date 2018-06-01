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


func HandleWX(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Body)
	fmt.Println(r.URL)
	if(r.Form==nil){
		fmt.Fprintf(w, "parameter need")
	}
	signature:= r.Form.Get("signature")
	timestamp:= r.Form.Get("timestamp")
	nonce := r.Form.Get("nonce")
	echostr := r.Form.Get("echostr")

	hashcode := makeSignature(timestamp,nonce)
	if hashcode == signature{
		fmt.Fprintf(w,echostr+"123")
	}else {
		fmt.Fprintf(w, "parameter error")
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