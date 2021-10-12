package test

import (
	"io"
	"net/http"
	"testing"
)

const (
	y = "\u2713"
	x = "\u2717"
)

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"https://www.ardanlabs.com/blog/index.xml", http.StatusOK,
		}, {
			"https://www.ardanlabs.com/blog/index.xml1", http.StatusNotFound,
		},
	}

	t.Log("given the need to test downloading different content")
	{
		for _, u := range urls {
			t.Logf("\twhen checking %s for status code %d", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("should be able to get the url", x, err)
				}

				t.Log("should be able to get the url", y)

				//goland:noinspection GoDeferInLoop
				defer func(Body io.ReadCloser) {
					_ = Body.Close()
				}(resp.Body)

				if resp.StatusCode == u.statusCode {
					t.Logf("should have a %d status. %v", u.statusCode, y)
				} else {
					t.Errorf("should have a %d status %v %v", u.statusCode, x, resp.StatusCode)
				}
			}
		}
	}
}
