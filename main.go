package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user"))
	})

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
