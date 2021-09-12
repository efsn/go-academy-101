package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Route() {
	http.HandleFunc("/send-json", SendJson)
}

func SendJson(rw http.ResponseWriter, _ *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Oho",
		Email: "oho@chain.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	err := json.NewEncoder(rw).Encode(&u)
	if err != nil {
		log.Fatalln(err)
	}
}
