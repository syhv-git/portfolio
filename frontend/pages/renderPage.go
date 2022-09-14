package pages

import (
	"html/template"
	"net/http"
)

var (
	dom = template.Must(template.ParseFiles("index.html"))
)

func render(w http.ResponseWriter, b []byte) {
	data := map[string]template.HTML{"Body": template.HTML(b)}
	if err := dom.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
