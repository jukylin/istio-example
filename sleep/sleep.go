package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sleep := r.Form.Get("sleep")
	sec,_ := strconv.Atoi(sleep)
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Fprint(w, "Sleep "+ sleep)
}

func retry(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	status := r.Form.Get("status")
	s ,_ := strconv.Atoi(status)

	w.WriteHeader(s)
	fmt.Fprint(w, "rep status" + status)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", sayhello)
	mux.HandleFunc("/retry", retry)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}