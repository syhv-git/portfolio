package tests

import (
	"bytes"
	"go-networking/frontend"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStaticFileServer(t *testing.T) {
	r := frontend.Routes("../assets")
	srv := httptest.NewServer(r)

	resp, err := http.Get(srv.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}

	cnt := resp.Header.Get("Content-Type")
	exp := "text/html; charset=utf-8"
	if cnt != exp {
		t.Errorf("** Failed: Content type of response was %v; expected %v", cnt, exp)
	}
}

func TestRouteHandlers(t *testing.T) {
	r := frontend.Routes("../assets")
	srv := httptest.NewServer(r)

	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()

	resp, err = http.Post(srv.URL+"/", "text/plain", bytes.NewReader([]byte("Hello Error!")))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}

	resp, err = http.Get(srv.URL + "/search-results")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}

	resp, err = http.Get(srv.URL + "/about")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}

	resp, err = http.Get(srv.URL + "/projects")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}

	resp, err = http.Get(srv.URL + "/resume")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("** Failed: Response from server was %v; expected %v", resp.StatusCode, http.StatusOK)
	}
}
