package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mdebelle/millionlogs/store"
)

func TestRouter(t *testing.T) {
	server := httptest.NewServer(handler())
	defer server.Close()
	type expected struct {
		status  int
		content string
	}

	path := "sample/hn_logs.tsv"
	store.Sample = &path

	testTable := []struct {
		Name     string
		url      string
		expected expected
	}{
		{"get count", "/1/queries/count/2015", expected{http.StatusOK, `{"count":573697}`}},
		{"get popular", "/1/queries/popular/2015?size=3", expected{http.StatusOK, `{"queries":[{"query":"http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice","count":6675},{"query":"http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045","count":4652},{"query":"http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1","count":3100}]}`}},
		{"not found", "/hello/world", expected{http.StatusNotFound, `404 page not found`}},
		{"get popular bad parameter", "/1/queries/popular/2015?size=hello", expected{http.StatusBadRequest, `not a number: hello`}},
		{"get count date bad parameter", "/1/queries/count/2015/2015", expected{http.StatusBadRequest, `"could not use '2015/2015' as a date parameter: bad format"`}},
	}

	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s%s", server.URL, testCase.url))
			if err != nil {
				t.Fatalf("could not perform get: %v", err)
			}
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could read all content: %v", err)
			}

			if resp.StatusCode != testCase.expected.status {
				t.Fatalf("expected status %d got %d", testCase.expected.status, resp.StatusCode)
			}
			if s := string(bytes.TrimSpace(content)); s != testCase.expected.content {
				t.Fatalf("expected content %s got %s", testCase.expected.content, s)
			}
		})
	}

	testTablePost := []struct {
		Name     string
		url      string
		expected expected
	}{
		{"post count", "/1/queries/count/2015", expected{http.StatusNotImplemented, `"Not Implemented"`}},
		{"post popular", "/1/queries/popular/2015?size=3", expected{http.StatusNotImplemented, `"Not Implemented"`}},
	}

	for _, testCase := range testTablePost {
		t.Run(testCase.Name, func(t *testing.T) {
			resp, err := http.Post(fmt.Sprintf("%s%s", server.URL, testCase.url), "", nil)
			if err != nil {
				t.Fatalf("could not perform post: %v", err)
			}
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could read all content: %v", err)
			}

			if resp.StatusCode != testCase.expected.status {
				t.Fatalf("expected status %d got %d", testCase.expected.status, resp.StatusCode)
			}
			if s := string(bytes.TrimSpace(content)); s != testCase.expected.content {
				t.Fatalf("expected content %s got %s", testCase.expected.content, s)
			}

		})
	}
}
