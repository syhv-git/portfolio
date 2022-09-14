package tests

import (
	"go-networking/frontend/pages"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLandingPageContent(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:9001", nil) //update url when deployed
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	h := http.HandlerFunc(pages.BuildLanding)
	h.ServeHTTP(rec, req)

	cnt := rec.Header().Get("Content-Type")
	exp := "text/html; charset=utf-8"
	if cnt != exp {
		t.Errorf("** Failed: Content type of response was %v; expected %v", cnt, exp)
	}
}
