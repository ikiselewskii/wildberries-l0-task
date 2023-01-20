package main

import "net/http"

func main(){
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":8080", nil)
}

func pong(rw http.ResponseWriter, r *http.Request){
	rw.WriteHeader(http.StatusAccepted)
	rw.Write([]byte("pong"))
}