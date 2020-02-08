package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Test",Test)
	http.HandleFunc("/Test1",Test1)
	http.ListenAndServe(":8080",nil)
}

func Test(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	date := val.Get("date")
	fmt.Println(date)
}

func Test1(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	date := val["date"][1]
	//date := val.Get("date")
	fmt.Println(date)
}
