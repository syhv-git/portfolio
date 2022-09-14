package pages

import (
	"net/http"
	"os"
)

func BuildLanding(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("assets/html/landing.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	render(w, b)
}
