package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/efsn/go-workshop-101/internal/handler"
)

const (
	Y = "\u2713"
	X = "\u2717"
)

func init() {
	handler.Route()
}

func TestSendJson(t *testing.T) {
	t.Log("given test the send json endpoint")
	{
		req, err := http.NewRequest("GET", "/send-json", nil)
		if err != nil {
			t.Fatal("should be able to create a request", X, err)
		}
		t.Log("should be able to create a request", Y)
		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("should receive 200", X, rw.Code)
		}
		t.Log("should receive200", Y)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("should decode the response", X)
		}
		t.Log("should decode the response", Y)

		if u.Name == "Oho" {
			t.Log("should have a Name", Y)
		} else {
			t.Error("should have a Name", X, u.Name)
		}

		if u.Email == "oho@chain.com" {
			t.Log("should have an Email", Y)
		} else {
			t.Error("should have an Email", X, u.Email)
		}
	}
}
