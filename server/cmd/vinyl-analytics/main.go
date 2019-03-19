package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("awoeighawoegh"))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}