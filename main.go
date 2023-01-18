package main

import "net/http"

func main(){
	http.HandleFunc("/ping", pong)
	http.ListenAndServe("8080", nil)
}

func pong(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(202)
}