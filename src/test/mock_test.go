package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

var feed = "{}"

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		_, err := fmt.Fprintln(w, feed)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownloadMock(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()

	t.Log("given the test downloading content")
	{
		t.Logf("when checking %s for status code %d", server.URL, statusCode)
		{
			res, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("should be able to make the get call", ballotX, err)
			}
			t.Log("should be able to make the get call", checkMark)

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(res.Body)

			if res.StatusCode != statusCode {
				t.Fatalf("should receive a %d status %v %v", statusCode, ballotX, res.StatusCode)
			}
			t.Logf("should receive a %d status %v", statusCode, checkMark)
		}
	}
}
