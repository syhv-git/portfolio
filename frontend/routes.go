package frontend

import (
	"github.com/gorilla/mux"
	"go-networking/frontend/pages"
	"net/http"
)

func Routes(dir string) *mux.Router {
	m := mux.NewRouter()
	m.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir)))).Methods("GET")

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.BuildPage(w, r, "landing")
	}).Methods("GET")
	m.HandleFunc("/search-results", pages.BuildSearch).Methods("GET")
	m.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		pages.BuildPage(w, r, "about")
	}).Methods("GET")
	m.HandleFunc("/projects", pages.BuildProjects).Methods("GET")
	m.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		pages.BuildPage(w, r, "resume")
	}).Methods("GET")

	return m
}
