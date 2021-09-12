package main

import (
	"go-workshop-101/src/handler"
	"log"
	"net/http"
)

func main() {
	handler.Route()
	log.Println("listener: started: listen on: 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
