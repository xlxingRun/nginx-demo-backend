package main

import (
	"fmt"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, 第一条消息")
	if err != nil {
		return
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, 第二条消息")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello", handler1)
	http.HandleFunc("/api/hello", handler2)
	err1 := http.ListenAndServe(":9000", nil)
	if err1 != nil {
		return
	}

}
