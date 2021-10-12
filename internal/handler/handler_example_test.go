package handler_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJson() {
	r, _ := http.NewRequest("GET", "/send-json", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("error: ", err)
	}

	fmt.Println(u)
	// Output:
	// {Oho oho@chain.com}
}
