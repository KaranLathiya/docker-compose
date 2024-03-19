package main

import "net/http"

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hi", hi)
	http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HELLO"))
}

func hi (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HI"))
}