package main

import (
	"fmt"
	"net/http"
	"log"
	"wechat_handler"
	"intercom_handler"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprint(w, "hello")
}


func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/wx", wechat_handler.HandleWX)
	http.HandleFunc("/intercom",intercom_handler.HandleIntercom)

	fmt.Println("listen on port 80")
	err := http.ListenAndServe(":80", nil);
	if  err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}