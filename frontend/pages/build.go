package pages

import (
	"html/template"
	"net/http"
	"os"
)

func BuildPage(w http.ResponseWriter, r *http.Request, p string) {
	b, err := os.ReadFile("assets/html/" + p + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render(w, b)
}

var (
	dom = template.Must(template.ParseFiles("index.html"))
)

func render(w http.ResponseWriter, b []byte) {
	n, err := os.ReadFile("assets/html/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f, err := os.ReadFile("assets/html/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]template.HTML{
		"Nav":  template.HTML(n),
		"Body": template.HTML(b),
		"Foot": template.HTML(f),
	}
	if err := dom.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
