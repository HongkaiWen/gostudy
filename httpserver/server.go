package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//chrome will redirect to the target url
	http.Redirect(w, r, "http://test.hongkai.cn:85/;jsessionid=3371D97462347B30C7D55ABD9C455E31", http.StatusFound)
	//chrome will canceled the request because the url is not valid
	//	http.Redirect(w, r, "http://test.hongkai.cn:85;jsessionid=3371D97462347B30C7D55ABD9C455E31", http.StatusFound)
}

func MainServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/", MainServer)
	err := http.ListenAndServe(":85", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
