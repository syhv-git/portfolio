package pages

import (
	"net/http"
	"os"
)

func BuildPage(w http.ResponseWriter, r *http.Request, p string) {
	b, err := os.ReadFile("assets/html/" + p + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	render(w, b)
}
