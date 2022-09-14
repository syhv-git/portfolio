package pages

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func BuildSearch(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := u.Query()
	query := param.Get("q")
	if query == "" {
		searchPage(w, r)
		return
	}
	page := param.Get("page")
	if page == "" {
		page = "1"
	}
	fmt.Println(query)

	b, err := os.ReadFile("assets/html/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	render(w, b)
}

func searchPage(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("assets/html/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	render(w, b)
}
