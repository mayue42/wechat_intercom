package main


import (
	"fmt"
	"net/http"
	"log"

)

const token string="token123"


func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.RequestURI)
	fmt.Println(r)
	fmt.Fprintf(w, "hello")
}

func HandleWX(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Body)
	//data = web.input()
	//if len(data) == 0:
	//return "hello, this is handle view"
	//signature = data.signature
	//timestamp = data.timestamp
	//nonce = data.nonce
	//echostr = data.echostr
	//token = "token123" #请按照公众平台官网\基本配置中信息填写
	//
	//list = [token, timestamp, nonce]
	//list.sort()
	//sha1 = hashlib.sha1()
	//map(sha1.update, list)
	//hashcode = sha1.hexdigest()
	//print "handle/GET func: hashcode, signature: ", hashcode, signature
	//if hashcode == signature:
	//return echostr
	fmt.Fprintf(w, "hello2")
}



func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/wx", HandleWX)
	fmt.Println("listen on port 9090")
	err := http.ListenAndServe(":9090", nil);
	if  err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}