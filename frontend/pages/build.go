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

func render(w http.ResponseWriter, b []byte) {
	dom := template.Must(template.ParseFiles("assets/html/index.html", "assets/html/navbar.html", "assets/html/footer.html"))
	data := map[string]template.HTML{
		"Body": template.HTML(b),
	}
	if err := dom.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
